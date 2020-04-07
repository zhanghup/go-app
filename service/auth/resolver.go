//go:generate go run github.com/99designs/gqlgen

package auth

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	c := Config{Resolvers: &Resolver{
		DB:  db,
		DBS: toolxorm.NewEngine(db),
	}}

	hu := handler.GraphQL(NewExecutableSchema(c))
	hu = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}(hu)

	return func(c *gin.Context) {
		hu.ServeHTTP(c.Writer, c.Request)
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
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
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
	return this.Token(ctx, *user.Id)
}

func (this mutationResolver) Token(ctx context.Context, uid string) (string, error) {
	token := new(beans.UserToken)
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {
		_, e := s.SF(`update {{ table "user_token" }} set status = 0 where uid = :uid and type = :type`, map[string]interface{}{
			"uid":  uid,
			"type": ty,
		}).Execute()
		if e != nil {
			return e
		}
		token.Id = tools.ObjectString()
		token.Status = tools.Ptr().Int(1)
		token.Uid = &uid
		token.Type = tools.Ptr().String(string(ty))
		token.Agent = tools.Ptr().String(this.me(ctx).GinContext().Request.UserAgent())
		token.Expire = tools.Ptr().Int64(2 * 60 * 60)
		token.Ops = tools.Ptr().Int64(0)
		_, e = s.Insert(token)
		if e != nil {
			return e
		}
		this.me(ctx).GinContext().SetCookie(directive.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
		return nil
	})

	if err != nil {
		return "", err
	}
	return *token.Id, nil
}

type queryResolver struct{ *Resolver }

func (queryResolver) Hello(ctx context.Context) (*string, error) {
	panic("implement me")
}
