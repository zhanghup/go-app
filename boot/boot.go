package boot

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/file"
	"net/http"
)

type BODB struct {
	Type    string `json:"type"`
	Uri     string `json:"uri"`
	ShowSql bool   `json:"show_sql"`
}

type BOWeb struct {
	Port string `json:"port"`
}

type BootOption struct {
	Version string `json:"version"`
	DB      BODB   `json:"db"`
	Web     BOWeb  `json:"web"`
}

func Boot(opt BootOption) {
	e, err := xorm.NewEngine(opt.DB.Type, opt.DB.Uri)
	if err != nil {
		panic(err)
	}
	app.Sync(e)
	initia.InitAction(e)
	e.ShowSQL(opt.DB.ShowSql)
	Router(e).Run(":" + opt.Web.Port)
}

func Router(e *xorm.Engine) *gin.Engine {
	g := gin.Default()
	// 文件操作
	{
		file.Gin(e, g)
	}

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
