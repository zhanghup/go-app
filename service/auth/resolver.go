//go:generate go run github.com/99designs/gqlgen

package auth

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"time"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	c := Config{Resolvers: &Resolver{
		DB:  db,
		DBS: txorm.NewEngine(db),
		Me:  directive.MyInfo,
	}}

	srv := handler.New(NewExecutableSchema(c))
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
	Me  func(g context.Context) directive.Me
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this mutationResolver) Logout(ctx context.Context) (bool, error) {
	tok := this.Me(ctx).GinContext().GetHeader(directive.GIN_AUTHORIZATION)

	if tok == "" {
		tokk, err := this.Me(ctx).GinContext().Cookie(directive.GIN_TOKEN)
		if err != nil {
			return false, err
		}
		tok = tokk
	}
	if tok == "" {
		return false, nil
	}
	_, err := this.DB.Table(beans.UserToken{}).Where("id = ?").Update(map[string]interface{}{"status": 0, "updated": time.Now().Unix()})
	return err == nil, err
}

func (this mutationResolver) LoginStatus(ctx context.Context) (bool, error) {
	tok := this.Me(ctx).GinContext().GetHeader(directive.GIN_AUTHORIZATION)

	if tok == "" {
		tokk, err := this.Me(ctx).GinContext().Cookie(directive.GIN_TOKEN)
		if err != nil {
			return false, err
		}
		tok = tokk
	}
	if tok == "" {
		return false, nil
	}
	t := beans.UserToken{}
	ok, err := this.DB.Where("id = ?", tok).Get(&t)
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
	if user.Salt == nil {
		if *user.Password == password {
			flag = true
		}
	} else {
		if *user.Password == tools.Crypto.Password(password, *user.Salt) {
			flag = true
		}
	}
	if !flag {
		return "", errors.New("用户名或者密码错误")
	}
	tok, err := this.Token(ctx, *user.Id, "pc")
	if err != nil {
		return "", err
	} else {
		go event.UserLogin("pc", &user)
	}
	return tok, nil
}

func (this mutationResolver) Token(ctx context.Context, uid, ty string) (string, error) {
	token := new(beans.UserToken)
	err := this.DBS.TS(func(sess *txorm.Session) error {
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

	return *token.Id, err
}

type queryResolver struct{ *Resolver }

func (q queryResolver) Hello(ctx context.Context) (*string, error) {
	panic("implement me")
}

func (this *Resolver) Query() QueryResolver {
	return &queryResolver{this}
}
