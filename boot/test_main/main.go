package main

import (
	"github.com/gin-gonic/gin"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/file"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
)

func main() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	cfg.InitConfig(box)
	e, err := txorm.NewXorm(cfg.DB)
	e.ShowSQL(true)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	//  同步表结构
	beans.Sync(e)

	// 初始化数据
	initia.InitAction(e)

	// http router
	err = tgin.NewGin(cfg.Web, func(g *gin.Engine) error {
		file.Gin(g.Group("/"), g.Group("/"), e)
		auth.Gin(g.Group("/"), e)
		api.Gin(g.Group("/"), e)
		return nil
	})

}
