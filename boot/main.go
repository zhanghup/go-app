package boot

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
	"os"
	"xorm.io/xorm"
)

type strujob struct {
	name string
	spec string
	fn   func(db *xorm.Engine) error
	flag []bool
}

type stru struct {
	name      string
	desc      string
	box       *rice.Box
	db        *xorm.Engine
	jobinit   bool
	jobs      []strujob
	initFns   []func(db *xorm.Engine)
	syncFns   []func(db *xorm.Engine)
	routerfns []func(g *gin.Engine, db *xorm.Engine)

	app *cli.App
}

type IBoot interface {
	Init(fns ...func(db *xorm.Engine)) IBoot
	SyncTables(fn ...func(db *xorm.Engine)) IBoot
	Jobs(name, spec string, fn func(db *xorm.Engine) error, flag ...bool) IBoot
	JobsMessageDealTimeout() IBoot
	Router(fn ...func(g *gin.Engine, db *xorm.Engine)) IBoot
	StartRouter()
}

/*
	box 配置文件位置
	name 服务名称
	desc 服务描述
*/
func Boot(box *rice.Box, name, desc string) IBoot {
	cfg.InitConfig(box)
	client := &stru{
		name:    name,
		desc:    desc,
		box:     box,
		jobinit: false,
	}

	// cli控制台命令
	{
		client.app = &cli.App{
			Name:  name,
			Usage: desc,
		}
	}

	// 是否需要初始化数据库
	{
		if client.db != nil {
			return client
		}
		e, err := txorm.NewXorm(cfg.DB)
		if err != nil {
			tog.Error(err.Error())
			panic(err)
		}
		client.db = e
		initia.InitDBTemplate(e)
	}

	return client
}

/*
	初始化默认的数据 - 必须初始化数据库
*/
func (this *stru) Init(fns ...func(db *xorm.Engine)) IBoot {
	this.initFns = append(this.initFns, fns...)
	return this
}

//  同步表结构
func (this *stru) SyncTables(fn ...func(db *xorm.Engine)) IBoot {
	this.syncFns = append(this.syncFns, fn...)
	return this
}

/*
	初始化定时任务 - 将与web服务一同运行
	@name: 任务名称
	@spec: * * * * * * cron表达式
	@fn: 任务执行方法
	@flag: 是否立即执行
*/
func (this *stru) Jobs(name, spec string, fn func(db *xorm.Engine) error, flag ...bool) IBoot {
	this.jobs = append(this.jobs, strujob{
		name: name,
		spec: spec,
		fn:   fn,
		flag: flag,
	})
	//if !this.jobinit {
	//	err := job.InitJobs(this.db)
	//	if err != nil {
	//		tog.Error(err.Error())
	//		panic(err)
	//	}
	//}
	//
	//err := job.AddJob(name, spec, fn, flag...)
	//if err != nil {
	//	tog.Error(err.Error())
	//	panic(err)
	//}
	return this
}

/*
	初始化内置的消息任务
*/
func (this *stru) JobsMessageDealTimeout() IBoot {
	this.Jobs("消息超时处理任务", "0 * * * * *", func(db *xorm.Engine) error {
		return ags.MessageTimeoutMark()
	})
	return this
}

// 自定义接口
func (this *stru) Router(fn ...func(g *gin.Engine, db *xorm.Engine)) IBoot {
	this.routerfns = append(this.routerfns, fn...)
	return this
}

func (this *stru) runWeb() error {
	return tgin.NewGin(cfg.Web, func(g *gin.Engine) error {
		for _, fn := range this.routerfns {
			fn(g, this.db)
		}
		return nil
	})
}

func (this *stru) StartRouter() {
	// 通知各个组件，数据库初始化已经完成
	event.XormDefaultInit(this.db)

	this.app.Commands = append(this.app.Commands, &cli.Command{
		Name:  "init",
		Usage: "初始化基础数据数据",
		Action: func(c *cli.Context) error {
			initia.InitDict(this.db)
			initia.InitUser(this.db)
			initia.InitMsgTemplate(this.db)
			if len(this.initFns) > 0 {
				for _, f := range this.initFns {
					f(this.db)
				}
			}
			return nil
		},
	})

	this.app.Commands = append(this.app.Commands, &cli.Command{
		Name:  "sync",
		Usage: "初始化基础数据数据",
		Action: func(c *cli.Context) error {
			beans.Sync(this.db)
			if len(this.syncFns) > 0 {
				for _, f := range this.syncFns {
					f(this.db)
				}
			}
			return nil
		},
	})

	this.app.Commands = append(this.app.Commands, &cli.Command{
		Name:  "run",
		Usage: "开启web服务",
		Action: func(c *cli.Context) error {
			return this.runWeb()
		},
	})

	// 开启任务
	this.app.Action = func(ctx *cli.Context) error {
		return this.runWeb()
	}
	err := this.app.Run(os.Args)
	if err != nil {
		tog.Error(err.Error())
	}
}
