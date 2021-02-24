//go:generate go run cmd/generator.go

package api

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools/tgql"
)

func NewResolver() *Resolver {
	return &Resolver{
		NewResolverTools(),
	}
}

func Gin(g gin.IRouter, sc ...graphql.ExecutableSchema) {
	s := source.NewExecutableSchema(source.Config{
		Resolvers: NewResolver(),
		Directives: source.DirectiveRoot{
			Perm: directive.Perm(gs.DB()),
			Root: directive.Root(gs.DB()),
		},
	})
	if len(sc) > 0 {
		s = sc[0]
	}
	ags.GinGql("/zpx/api", g.Group("/", directive.WebAuth()), s)
}

type Resolver struct {
	*ResolverTools
}

type ResolverTools struct {
	Loader func(ctx context.Context) tgql.Loader
	Me     func(ctx context.Context) *ca.User
}

func NewResolverTools() *ResolverTools {
	return &ResolverTools{
		Loader: tgql.DataLoaden,
		Me:     directive.MyInfo,
	}
}
