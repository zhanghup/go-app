package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/api"
	"github.com/zhanghup/go-app/api/gs"
)

func main() {
	g := gin.Default()
	api.Playground(g, "/query")

	c := gs.Config{Resolvers: &api.Resolver{}}
	hu := handler.GraphQL(gs.NewExecutableSchema(c))
	g.POST("/query", func(c *gin.Context) {
		c.Request.Header.Set("Content-Type", "application/json")
		hu.ServeHTTP(c.Writer, c.Request)
	})
	g.GET("/qq", func(c *gin.Context) {
		handler.Playground("标题", "/query").ServeHTTP(c.Writer, c.Request)
	})

	g.Run(":8899")
}
