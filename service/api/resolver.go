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
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/gs"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func NewResolver(db *xorm.Engine) lib.Config {

	//cache, err := directive.NewDictCache(db)
	//if err != nil {
	//	panic(err)
	//}

	resolver := &Resolver{
		DB:     db,
		DBS:    txorm.NewEngine(db),
		Loader: tgql.DataLoaden,
		Me:     directive.MyInfo,
	}

	return lib.Config{
		Resolvers: resolver,
		Directives: lib.DirectiveRoot{
			Perm: directive.Perm(),
		},
	}
}

func ggin(db *xorm.Engine) func(c *gin.Context) {
	srv := handler.New(lib.NewExecutableSchema(NewResolver(db)))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
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

	g.Group("/", directive.WebAuth(db)).POST("/api", ggin(db))
	g.Group("/", directive.WebAuth(db)).GET("/api", ggin(db))
	gs.Playground(g, "/api/playground1", "/api")
	g.GET("/api/playground2", func(c *gin.Context) {
		playground.Handler("标题", "/api")(c.Writer, c.Request)
	})
}

type Resolver struct {
	DB        *xorm.Engine
	DBS       *txorm.Engine
	Loader    func(ctx context.Context) tgql.Loader
	Me        func(ctx context.Context) directive.Me
	DictCache func(dict string) (*beans.Dict, []beans.DictItem, bool)
}
