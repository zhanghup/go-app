//go:generate go run github.com/99designs/gqlgen

package auth

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools"
	"net/http"
)

func ggin(e *xorm.Engine) func(c *gin.Context) {
	c := Config{Resolvers: &Resolver{
		DB:     e,
		Loader: gs.DataLoaden,
	}}

	hu := handler.GraphQL(NewExecutableSchema(c))
	hu = gs.DataLoadenMiddleware(e, hu)
	hu = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}(hu)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), gs.GIN_CONTEXT, c)
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}


func Gin(e *xorm.Engine, g *gin.Engine) {
	g.POST("/auth", ggin(e))
	api.Playground(g, "/auth/playground1", "/auth")
	g.GET("/auth/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/auth").ServeHTTP(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB     *xorm.Engine
	Loader func(ctx context.Context) gs.Loader
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this mutationResolver) Login(ctx context.Context, account string, password string) (string, error) {
	user := app.User{}
	ok, err := this.DB.Where("account = ? and status = 1", account).Get(&user)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", gin.NewErr("用户不存在")
	}

	flag := false
	if user.Slat == nil {
		if *user.Password == password {
			flag = true
		}
	} else {
		if *user.Password == tools.Password(password, *user.Slat) {
			flag = true
		}
	}
	if !flag {
		return "", gin.NewErr("登录失败")
	}
	return this.Token(ctx, *user.Id, TokenPc)
}

func (this mutationResolver) Token(ctx context.Context, uid string, ty TokenType) (string, error) {
	token := new(app.UserToken)
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {
		_, e := s.SF(`update {{ table "user_token" }} set status = 0 where user = :user and type = :type`, map[string]interface{}{
			"user": uid,
			"type": ty,
		}).Execute()
		if e != nil {
			return e
		}
		token.Id = tools.ObjectString()
		token.Status = tools.Ptr().Int(1)
		token.User = &uid
		token.Type = tools.Ptr().String(string(ty))
		token.Agent = tools.Ptr().String(gs.GinContext(ctx).Request.UserAgent())
		token.Expire = tools.Ptr().Int64(2 * 60 * 60)
		token.Ops = tools.Ptr().Int64(0)
		_, e = s.Insert(token)
		if e != nil {
			return e
		}
		gs.GinContext(ctx).SetCookie(gs.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
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
