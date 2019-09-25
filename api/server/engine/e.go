package engine

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/auth"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-app/api"
)

func Router() *gin.Engine {
	g := gin.Default()

	e, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	app.Sync(e)
	e.ShowSQL(true)

	g.POST("/base", api.Gin(e))
	g.POST("/auth", auth.Gin(e))
	api.Playground(g, "/base/playground1","/base")
	api.Playground(g, "/auth/playground1","/auth")

	g.GET("/base/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/base").ServeHTTP(c.Writer, c.Request)
	})
	g.GET("/auth/playground2", func(c *gin.Context) {
		handler.Playground("标题", "/auth").ServeHTTP(c.Writer, c.Request)
	})
	return g
}
