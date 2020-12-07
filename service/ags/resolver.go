//go:generate go run github.com/99designs/gqlgen

package ags

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhanghup/go-app/service/ags/resolvers"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools/tgql"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func gqlschemaFmt(db *xorm.Engine, schema graphql.ExecutableSchema) func(c *gin.Context) {
	srv := handler.New(schema)
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

func Gql(gqlpath string, gqlrouter gin.IRouter, gqlSchema graphql.ExecutableSchema, db *xorm.Engine) {
	gqlrouter.POST(gqlpath, gqlschemaFmt(db, gqlSchema))
	gqlrouter.GET(gqlpath, gqlschemaFmt(db, gqlSchema))
	Playground(gqlrouter, gqlpath+"/playground1", gqlpath)
	gqlrouter.GET(gqlpath+"/playground2", func(c *gin.Context) {
		playground.Handler("标题", gqlpath)(c.Writer, c.Request)
	})
}

func Gin(auth, any gin.IRouter, db *xorm.Engine) {
	Gql("/zpx/ags", any, source.NewExecutableSchema(source.Config{Resolvers: resolvers.NewResolver(db)}), db)
	NewUploader(db).GinRouter(auth.Group("/zpx/agx"), any.Group("/zpx/ags"))

}
