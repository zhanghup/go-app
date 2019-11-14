package boot

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/auth"
	"net/http"
)

func Boot() {
	e, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	app.Sync(e)
	initia.InitDict(e)
	initia.InitUser(e)
	e.ShowSQL(true)
	Router(e).Run(":8899")
}

func Router(e *xorm.Engine) *gin.Engine {
	g := gin.Default()

	// http服务状态监听
	{
		g.Use(StatsRequest())
		g.GET("/stats", func(c *gin.Context) {
			c.JSON(http.StatusOK, StatsReport())
		})
	}

	// 基础数据路由
	{
		g.POST("/base", api.Gin(e))
		api.Playground(g, "/base/playground1", "/base")
		g.GET("/base/playground2", func(c *gin.Context) {
			handler.Playground("标题", "/base").ServeHTTP(c.Writer, c.Request)
		})
	}

	// 授权路由
	{
		g.POST("/auth", auth.Gin(e))
		api.Playground(g, "/auth/playground1", "/auth")
		g.GET("/auth/playground2", func(c *gin.Context) {
			handler.Playground("标题", "/auth").ServeHTTP(c.Writer, c.Request)
		})
	}

	return g
}
