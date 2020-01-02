package boot

import (
	rice "github.com/giter/go.rice"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/ctx"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/file"
)

func Boot(fn func() (*rice.Box, error)) {
	// 初始化配置文件
	ctx.InitConfFile()
	box, err := fn()
	if err != nil {
		panic(err)
	}
	ctx.InitConfig(box)

	if ctx.DBEnable() {
		// 同步表结构
		beans.Sync()

		// 初始化表数据
		initia.InitAction()
	}

	// 初始化基础路由
	if ctx.WebEnable() {
		// http服务状态监听
		{
			ctx.Web().Engine().Use(api.StatsRequest())
		}

		// 文件操作
		{
			file.Gin()
		}

		// 基础数据路由
		{
			api.Gin()
		}

		// 授权路由
		{
			auth.Gin()
		}
	}
}

func Run() error {
	if !ctx.WebEnable() {
		return nil
	}
	return ctx.Web().Run()
}
