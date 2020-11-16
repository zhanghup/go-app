//go:generate go run github.com/99designs/gqlgen

package auth

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/auth/source"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	c := source.Config{Resolvers: &Resolver{
		DB:  db,
		DBS: txorm.NewEngine(db),
		Gin: func(g context.Context) *gin.Context {
			gg := g.Value(directive.GIN_CONTEXT)
			ggg := gg.(*gin.Context)
			return ggg
		},
	}}

	srv := handler.New(source.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		ctx := context.WithValue(c.Request.Context(), directive.GIN_CONTEXT, c)
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(g gin.IRouter, db *xorm.Engine) {
	g.POST("/auth", ggin(db))
	gs.Playground(g, "/auth/playground1", "/auth")
	g.GET("/auth/playground2", func(c *gin.Context) {
		playground.Handler("标题", "/auth")(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB  *xorm.Engine
	DBS *txorm.Engine
	Gin func(g context.Context) *gin.Context
}

func (r *Resolver) Mutation() source.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this *mutationResolver) Logout(ctx context.Context) (bool, error) {
	tok := this.Gin(ctx).GetHeader(directive.GIN_AUTHORIZATION)

	if tok == "" {
		tokk, err := this.Gin(ctx).Cookie(directive.GIN_TOKEN)
		if err != nil {
			return false, err
		}
		tok = tokk
	}
	if tok == "" {
		return false, nil
	}
	_, err := this.DB.Table(beans.Token{}).Where("id = ?", tok).Update(map[string]interface{}{"status": 0})
	return err == nil, err
}

func (this *queryResolver) LoginStatus(ctx context.Context) (bool, error) {
	tok := this.Gin(ctx).GetHeader(directive.GIN_AUTHORIZATION)

	if tok == "" {
		tokk, err := this.Gin(ctx).Cookie(directive.GIN_TOKEN)
		if err != nil && err != http.ErrNoCookie {
			return false, err
		}
		tok = tokk
	}
	if tok == "" {
		return false, nil
	}
	t := beans.Token{}
	ok, err := this.DB.Where("id = ? and status = 1", tok).Get(&t)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	if time.Now().Unix() > *t.Updated+*t.Expire {
		return false, nil
	}
	return true, nil
}

func (this *mutationResolver) Login(ctx context.Context, username string, password string) (string, error) {
	acc := beans.Account{}
	// 1. 找到账户
	ok, err := this.DB.Where("username = ? and status = '1' and type = 'password'", username).Get(&acc)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("账户不存在")
	}

	// 2. 验证密码
	flag := false
	if acc.Salt == nil {
		if *acc.Password == password {
			flag = true
		}
	} else {
		if *acc.Password == tools.Crypto.Password(password, *acc.Salt) {
			flag = true
		}
	}
	if !flag {
		return "", errors.New("用户名或者密码错误")
	}

	// 3. 找到用户
	user := beans.User{}
	ok, err = this.DB.Where("id = ? and status = '1'", acc.Uid).Get(&user)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("用户不存在")
	}

	tok, err := this.Token(ctx, *user.Id, *acc.Id)
	if err != nil {
		return "", err
	} else {
		go event.UserLogin(acc, user)
	}
	return tok, nil
}

func (this *mutationResolver) Token(ctx context.Context, uid, aid string) (string, error) {
	token := new(beans.Token)
	err := this.DBS.TS(func(sess *txorm.Session) error {
		e := sess.SF(`update token set status = '0' where uid = :uid and aid = :aid`, map[string]interface{}{
			"uid": uid,
			"aid": aid,
		}).Exec()
		if e != nil {
			return e
		}
		token.Id = tools.Ptr.Uid()
		token.Status = tools.Ptr.String("1")
		token.Uid = &uid
		token.Aid = &aid
		token.Agent = tools.Ptr.String(this.Gin(ctx).Request.UserAgent())
		token.Expire = tools.Ptr.Int64(2 * 60 * 60)
		token.Ops = tools.Ptr.Int64(0)
		e = sess.Insert(token)
		if e != nil {
			return e
		}
		this.Gin(ctx).SetCookie(directive.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
		return nil

	})

	return *token.Id, err
}

type queryResolver struct{ *Resolver }

func (q *queryResolver) Hello(ctx context.Context) (*string, error) {
	return tools.Ptr.String("hello world"), nil
}

func (this *Resolver) Query() source.QueryResolver {
	return &queryResolver{this}
}
