package boot

import (
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
	initia.InitAction(e)
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
		api.Gin(e, g)
	}

	// 授权路由
	{
		auth.Gin(e, g)
	}

	return g
}
