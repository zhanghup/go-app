//go:generate go run cmd/generator.go

package awxmp

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/awxmp/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"github.com/zhanghup/go-tools/tog"
	"github.com/zhanghup/go-tools/wx/wxmp"
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
		},
	})
	if len(sc) > 0 {
		s = sc[0]
	}
	ags.GinGql("/zpx/wxmp", g.Group("/", directive.WxmpAuth(ags.DefaultDB())), s, ags.DefaultDB())
	g.POST("/zpx/wxmp/pay/callback", PayCallback)
}

type Resolver struct {
	*ResolverTools
}

type ResolverTools struct {
	DBS    func(ctx context.Context) txorm.ISession
	Sess   func(ctx context.Context) txorm.ISession
	Loader func(ctx context.Context) tgql.Loader
	Wxme   func(ctx context.Context) *ca.WxmpUser
	Wxmp   wxmp.IEngine
}

func NewResolverTools() *ResolverTools {
	dbs := txorm.NewEngine(ags.DefaultDB())
	return &ResolverTools{
		DBS: func(ctx context.Context) txorm.ISession {
			return dbs.NewSession(true, ctx)
		},
		Sess: func(ctx context.Context) txorm.ISession {
			sess := dbs.Session(ctx)
			err := sess.Begin()
			if err != nil {
				tog.Error("【开启事务异常！！！】")
			}
			return sess
		},
		Loader: tgql.DataLoaden,
		Wxme:   directive.MyWxmpUser,
		Wxmp:   wxmp.NewEngine(&cfg.Wxmp),
	}
}
