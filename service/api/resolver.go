//go:generate go run cmd/generator.go

package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

func NewResolver(db *xorm.Engine) *Resolver {
	return &Resolver{
		NewResolverTools(db),
	}
}

func Gin(g gin.IRouter, db *xorm.Engine) {
	config := source.Config{
		Resolvers: NewResolver(db),
		Directives: source.DirectiveRoot{
			Perm: directive.Perm(db),
			Root: directive.Root(db),
		},
	}
	ags.GinGql("/zpx/api", g.Group("/", directive.WebAuth(db)), source.NewExecutableSchema(config), db)
}

type Resolver struct {
	*ResolverTools
}

type ResolverTools struct {
	DBS    func(ctx context.Context) txorm.ISession
	Sess   func(ctx context.Context) txorm.ISession
	Loader func(ctx context.Context) tgql.Loader
	Me     func(ctx context.Context) *ca.User
}

func NewResolverTools(db *xorm.Engine) *ResolverTools {
	dbs := txorm.NewEngine(db)
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
		Me:     directive.MyInfo,
	}
}
