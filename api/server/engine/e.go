package engine

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-app/api"
)

func Router() *gin.Engine {
	g := gin.Default()
	api.Playground(g, "/query")

	//e, err := xorm.NewEngine("sqlite3", "./test.db")
	e, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	app.Sync(e)
	e.ShowSQL(true)

	g.POST("/query", api.Gin(e))
	g.GET("/qq", func(c *gin.Context) {
		handler.Playground("标题", "/query").ServeHTTP(c.Writer, c.Request)
	})
	return g
}
