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
	"github.com/zhanghup/go-app/service/gin_stat"
)

func Boot(box *rice.Box) {
	// 初始化配置文件
	cfg.InitConfig(box)

	// 同步表结构
	beans.Sync()

	// 初始化表数据
	initia.InitAction()

	// 初始化基础路由
	{
		// http服务状态监听
		{
			gin_stat.Gin()
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
	return cfg.Web().Run()
}
