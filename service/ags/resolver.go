//go:generate go run github.com/99designs/gqlgen

package ags

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/ags/resolvers"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/directive"
	"xorm.io/xorm"
)

func ggin(db *xorm.Engine) func(c *gin.Context) {
	c := source.Config{Resolvers: resolvers.NewResolver(db)}

	srv := handler.New(source.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		ctx := context.WithValue(c.Request.Context(), directive.GIN_CONTEXT, c)
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func Gin(g gin.IRouter, db *xorm.Engine) {
	g.POST("/ags", ggin(db))
	Playground(g, "/ags/playground1", "/ags")
	g.GET("/ags/playground2", func(c *gin.Context) {
		playground.Handler("标题", "/ags")(c.Writer, c.Request)
	})
}
