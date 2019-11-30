package boot

import (
	rice "github.com/giter/go.rice"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/file"
)

func Boot(fn func() (*rice.Box, error)) {
	// 初始化配置文件
	cfg.InitConfFile()
	box, err := fn()
	if err != nil {
		panic(err)
	}
	cfg.InitConfig(box)

	if cfg.DBEnable() {
		// 同步表结构
		beans.Sync()

		// 初始化表数据
		initia.InitAction()
	}

	// 初始化基础路由
	if cfg.WebEnable() {
		// http服务状态监听
		{
			cfg.Web().Engine().Use(api.StatsRequest())
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
	if !cfg.WebEnable() {
		return nil
	}
	return cfg.Web().Run()
}
