//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-app/service/loaders"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-wxmp"
	"net/http"
)

func ggin() func(c *gin.Context) {
	resolver := &Resolver{
		DB:     cfg.DB().Engine(),
		Loader: loaders.DataLoaden,
		my:     directive.MewMe,
	}

	if cfg.WxmpEnable() {
		resolver.wxmp = wxmp.NewContext(cfg.Wxmp().Appid, cfg.Wxmp().AppSecret, cfg.Wxmp().Token)
	}

	c := lib.Config{
		Resolvers: resolver,
		Directives: lib.DirectiveRoot{
			Perm: directive.Perm(),
		},
	}

	hu := handler.GraphQL(lib.NewExecutableSchema(c))
	hu = loaders.DataLoadenMiddleware(cfg.DB().Engine(), hu)
	hu = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}(hu)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), directive.GIN_CONTEXT, c)
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin() {
	cfg.Web().Engine().Group("/", directive.UserAuth()).POST("/api", ggin())
	gs.Playground("/api/playground1", "/api")
	cfg.Web().Engine().GET("/api/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/api").ServeHTTP(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB     *xorm.Engine
	Loader func(ctx context.Context) loaders.Loader
	my     func(ctx context.Context) directive.Me
	wxmp   wxmp.IContext
}

func (r *Resolver) Mutation() lib.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() lib.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }



func (this mutationResolver) World(ctx context.Context) (*string, error) {
	return tools.Ptr().String("hello"), nil
}

type queryResolver struct{ *Resolver }



func (this queryResolver) Hello(ctx context.Context) (*string, error) {
	return tools.Ptr().String("world"), nil
}
