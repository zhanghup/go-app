package main

import (
	"github.com/gin-gonic/gin"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/file"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"github.com/zhanghup/go-tools/toolgin"
)

func main() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	cfg.InitConfig(box)
	e, err := toolxorm.NewXorm(cfg.DB)
	if err != nil {
		panic(err)
	}
	//  同步表结构
	beans.Sync(e)

	// 初始化数据
	initia.InitAction(e)

	// http router
	err = toolgin.NewGin(cfg.Web, func(g *gin.Engine) error {
		file.Gin(g.Group("/"), g.Group("/"), e)
		auth.Gin(g.Group("/auth"), e)
		return nil
	})

}
