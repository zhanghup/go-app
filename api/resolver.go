//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app/api/gs"
	"github.com/zhanghup/go-tools"
	"net/http"
)

func Gin(e *xorm.Engine, nexts ...http.HandlerFunc) func(c *gin.Context) {
	c := gs.Config{Resolvers: &Resolver{
		DB:     e,
		Loader: gs.DataLoaden,
	}}

	hu := handler.GraphQL(gs.NewExecutableSchema(c))
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
		hu.ServeHTTP(c.Writer, c.Request)
	}
}

type Resolver struct {
	DB     *xorm.Engine
	Loader func(ctx context.Context) gs.Loader
}

func (r *Resolver) Mutation() gs.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gs.QueryResolver {
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
