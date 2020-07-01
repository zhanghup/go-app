//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	resolver := &Resolver{
		DB:     db,
		DBS:    txorm.NewEngine(db),
		Loader: tgql.DataLoaden,
		my:     directive.MyInfo,
	}
	c := lib.Config{
		Resolvers: resolver,
		Directives: lib.DirectiveRoot{
			Perm: directive.Perm(),
		},
	}

	srv := handler.New(lib.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	hu := tgql.DataLoadenMiddleware(db, srv)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), directive.GIN_CONTEXT, c)
		c.Header("Content-Type", "application/json")
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(g gin.IRouter, db *xorm.Engine) {
	g.Group("/", directive.WebAuth(db)).POST("/api", ggin(db))
	g.Group("/", directive.WebAuth(db)).GET("/api", ggin(db))
	gs.Playground(g, "/api/playground1", "/api")
	g.GET("/api/playground2", func(c *gin.Context) {
		playground.Handler("标题", "/api")(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB     *xorm.Engine
	DBS    *txorm.Engine
	Loader func(ctx context.Context) tgql.Loader
	my     func(ctx context.Context) directive.Me
}
