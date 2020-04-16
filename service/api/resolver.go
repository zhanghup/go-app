//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"net/http"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	resolver := &Resolver{
		DB:     db,
		DBS:    txorm.NewEngine(db),
		Loader: tgql.DataLoaden,
	}
	c := lib.Config{
		Resolvers: resolver,
		Directives: lib.DirectiveRoot{
			Perm: directive.Perm(),
		},
	}

	hu := handler.GraphQL(lib.NewExecutableSchema(c))
	hu = tgql.DataLoadenMiddleware(db, hu)
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

func Gin(g gin.IRouter, db *xorm.Engine) {
	g.Group("/", directive.WebAuth(db)).POST("/api", ggin(db))
	gs.Playground(g, "/api/playground1", "/api")
	g.GET("/api/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/api/api").ServeHTTP(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB     *xorm.Engine
	DBS    *txorm.Engine
	Loader func(ctx context.Context) tgql.Loader
	my     func(ctx context.Context) directive.Me
}

func (r *Resolver) Mutation() lib.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() lib.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this mutationResolver) World(ctx context.Context) (*string, error) {
	panic("implement me")
}

type queryResolver struct{ *Resolver }

func (this queryResolver) Hello(ctx context.Context) (*string, error) {
	panic("implement me")
}
