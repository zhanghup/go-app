//go:generate go run github.com/99designs/gqlgen

package ags

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhanghup/go-app/service/ags/resolvers"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/directive"
	"net/http"
	"time"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	c := source.Config{Resolvers: resolvers.NewResolver(db)}

	srv := handler.New(source.NewExecutableSchema(c))
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

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		ctx := context.WithValue(c.Request.Context(), directive.GIN_CONTEXT, c)
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(auth, any gin.IRouter, db *xorm.Engine) {
	any.POST("/ags", ggin(db))
	any.GET("/ags", ggin(db))
	Playground(any, "/ags/playground1", "/ags")
	any.GET("/ags/playground2", func(c *gin.Context) {
		playground.Handler("标题", "/ags")(c.Writer, c.Request)
	})

	up := NewUploader(db)
	auth.POST("/ags/upload", up.Upload())
	any.GET("/ags/upload/:id", up.Get())
	any.GET("/ags/upload/:id/:width/:height", up.Resize())
	any.GET("/ags/upload/:id/:width", up.Resize())
}
