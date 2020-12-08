//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api/source"
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
		},
	}
	ags.Gql("/zpx/api", g.Group("/", directive.WebAuth(db)), source.NewExecutableSchema(config), db)
}

type Resolver struct {
	*ResolverTools
}

type ResolverTools struct {
	DB        *xorm.Engine
	DBS       func() *txorm.Engine
	Sess      func(ctx context.Context) txorm.ISession
	Loader    func(ctx context.Context) tgql.Loader
	Me        func(ctx context.Context) directive.Me
	DictCache func(dict string) (*beans.Dict, []beans.DictItem, bool)
}

func NewResolverTools(db *xorm.Engine) *ResolverTools {
	dbs := txorm.NewEngine(db)
	return &ResolverTools{
		DB: db,
		DBS: func() *txorm.Engine {
			return dbs
		},
		Sess: func(ctx context.Context) txorm.ISession {
			sess := dbs.NewSession(ctx)
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
