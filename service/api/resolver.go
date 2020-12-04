//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func NewResolver(db *xorm.Engine) *Resolver {

	return &Resolver{
		DB:     db,
		DBS:    txorm.NewEngine(db),
		Loader: tgql.DataLoaden,
		Me:     directive.MyInfo,
	}
}

func ggin(db *xorm.Engine) func(c *gin.Context) {
	config := source.Config{
		Resolvers: 	NewResolver(db),
		Directives: source.DirectiveRoot{
			Perm: directive.Perm(db),
		},
	}
	srv := handler.New(source.NewExecutableSchema(config))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {

			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			gtx := ctx.Value(directive.GIN_CONTEXT).(*gin.Context)
			_, err := directive.WebAuthFunc(db, gtx)
			return ctx, err
		},
	})
	srv.Use(extension.Introspection{})

	hu := tgql.DataLoadenMiddleware(db, srv)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), directive.GIN_CONTEXT, c)
		c.Header("Content-Type", "application/json")
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(g gin.IRouter, db *xorm.Engine) {

	g.Group("/", directive.WebAuth(db)).POST("/zpx/api", ggin(db))
	g.Group("/", directive.WebAuth(db)).GET("/zpx/api", ggin(db))
	ags.Playground(g, "/zpx/api/playground1", "/zpx/api")
	g.GET("/zpx/api/playground2", func(c *gin.Context) {
		playground.Handler("标题", "/zpx/api")(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB        *xorm.Engine
	DBS       *txorm.Engine
	Loader    func(ctx context.Context) tgql.Loader
	Me        func(ctx context.Context) directive.Me
	DictCache func(dict string) (*beans.Dict, []beans.DictItem, bool)
}
