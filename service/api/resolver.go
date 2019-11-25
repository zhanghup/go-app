//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools"
	"net/http"
)

func ggin(e *xorm.Engine, nexts ...http.HandlerFunc) func(c *gin.Context) {
	c := lib.Config{
		Resolvers: &Resolver{
			DB:     e,
			Loader: gs.DataLoaden,
		},
		Directives: lib.DirectiveRoot{
			Perm: gs.Perm(),
		},
	}

	hu := handler.GraphQL(lib.NewExecutableSchema(c))
	hu = gs.DataLoadenMiddleware(e, hu)
	hu = func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}(hu)

	for _, obj := range nexts {
		hu = func(next http.HandlerFunc) http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		}(obj)
	}

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), gs.GIN_CONTEXT, c)
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(e *xorm.Engine, g *gin.Engine) {
	g.Group("/", userAuth(e)).POST("/base", ggin(e))
	gs.Playground(g, "/base/playground1", "/base")
	g.GET("/base/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/base").ServeHTTP(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB     *xorm.Engine
	Loader func(ctx context.Context) gs.Loader
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
