//go:generate go run github.com/99designs/gqlgen

package auth

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"net/http"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	c := Config{Resolvers: &Resolver{
		DB:  db,
		DBS: toolxorm.NewEngine(db),
		Me:  directive.MyInfo,
	}}

	hu := handler.GraphQL(NewExecutableSchema(c))
	hu = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}(hu)

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, directive.GIN_CONTEXT, c)
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(g gin.IRouter, db *xorm.Engine) {
	g.POST("/auth", ggin(db))
	gs.Playground(g, "/auth/playground1", "/auth")
	g.GET("/auth/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/auth").ServeHTTP(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB  *xorm.Engine
	DBS *toolxorm.Engine
	Me  func(g context.Context) directive.Me
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this mutationResolver) Login(ctx context.Context, account string, password string) (string, error) {
	user := beans.User{}
	ok, err := this.DB.Where("account = ? and status = 1", account).Get(&user)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("用户不存在")
	}

	flag := false
	if user.Slat == nil {
		if *user.Password == password {
			flag = true
		}
	} else {
		if *user.Password == tools.Crypto.Password(password, *user.Slat) {
			flag = true
		}
	}
	if !flag {
		return "", errors.New("用户名或者密码错误")
	}
	return this.Token(ctx, *user.Id, "pc")
}

func (this mutationResolver) Token(ctx context.Context, uid, ty string) (string, error) {
	token := new(beans.UserToken)
	this.DBS.TS(func(sess *toolxorm.Session) error {
		e := sess.SF(`update user_token set status = 0 where uid = :uid and type = :type`, map[string]interface{}{
			"uid":  uid,
			"type": ty,
		}).Exec()
		if e != nil {
			return e
		}
		token.Id = tools.Ptr.Uid()
		token.Status = tools.Ptr.Int(1)
		token.Uid = &uid
		token.Type = tools.Ptr.String(string(ty))
		token.Agent = tools.Ptr.String(this.Me(ctx).GinContext().Request.UserAgent())
		token.Expire = tools.Ptr.Int64(2 * 60 * 60)
		token.Ops = tools.Ptr.Int64(0)
		e = sess.Insert(token)
		if e != nil {
			return e
		}
		this.Me(ctx).GinContext().SetCookie(directive.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
		return nil

	})

	return *token.Id, nil
}
