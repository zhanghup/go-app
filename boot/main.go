package boot

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/job"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
	"os"
)

type strujob struct {
	name string
	spec string
	fn   func() error
	flag []bool
}

type stru struct {
	name      string
	desc      string
	box       *rice.Box
	jobinit   bool
	jobs      []strujob
	initFns   []func()
	syncFns   []func()
	routerfns []func(g *gin.Engine)
	commonds  []*cli.Command

	app *cli.App
}

type IBoot interface {
	Cmd(fns ...func() []cli.Command) IBoot
	Init(fns ...func()) IBoot
	SyncTables(fn ...func()) IBoot
	Jobs(name, spec string, fn func() error, flag ...bool) IBoot
	Router(fn ...func(g *gin.Engine)) IBoot
	StartRouter()
}

/*
	box 配置文件位置
	name 服务名称
	desc 服务描述
*/
func Boot(box *rice.Box, name, desc string) IBoot {
	gs.ConfigInit(box)
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
		e, err := txorm.NewXorm(gs.Config.Database)
		if err != nil {
			tog.Error(err.Error())
			panic(err)
		}
		gs.Init(e)
		initia.InitDBTemplate()
	}

	return client
}

/* 初始化默认的数据 - 必须初始化数据库 */
func (this *stru) Init(fns ...func()) IBoot {
	this.initFns = append(this.initFns, fns...)
	return this
}

//  同步表结构
func (this *stru) SyncTables(fn ...func()) IBoot {
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
func (this *stru) Jobs(name, spec string, fn func() error, flag ...bool) IBoot {
	this.jobs = append(this.jobs, strujob{
		name: name,
		spec: spec,
		fn:   fn,
		flag: flag,
	})
	return this
}

// 自定义接口
func (this *stru) Router(fn ...func(g *gin.Engine)) IBoot {
	this.routerfns = append(this.routerfns, fn...)
	return this
}

// 自定义命令
func (this *stru) Cmd(fns ...func() []cli.Command) IBoot {
	for i := range fns {
		cmds := fns[i]()
		for j := range cmds {
			this.commonds = append(this.commonds, &cmds[j])
		}
	}
	return this
}

func (this *stru) runWeb() error {
	return tgin.NewGin(gs.Config.Web, func(g *gin.Engine) error {
		// 开启定时任务
		if len(this.jobs) > 0 {
			err := job.InitJobs()
			if err != nil {
				panic(err)
			}
			for _, j := range this.jobs {
				err := job.AddJob(j.name, j.spec, j.fn, j.flag...)
				if err != nil {
					panic(err)
				}
			}

		}

		// 初始化路由
		for _, fn := range this.routerfns {
			fn(g)
		}
		return nil
	})
}

func (this *stru) StartRouter() {
	this.app.Commands = append(this.app.Commands, &cli.Command{
		Name:  "init",
		Usage: "初始化基础数据数据",
		Action: func(c *cli.Context) error {
			initia.InitDict()
			initia.InitUser()
			initia.InitMsgTemplate()
			if len(this.initFns) > 0 {
				for _, f := range this.initFns {
					f()
				}
			}
			return nil
		},
	})

	this.app.Commands = append(this.app.Commands, &cli.Command{
		Name:  "sync",
		Usage: "同步表结构到数据库",
		Action: func(c *cli.Context) error {
			beans.Sync(gs.DB())
			if len(this.syncFns) > 0 {
				for _, f := range this.syncFns {
					f()
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
	this.app.Commands = append(this.app.Commands, this.commonds...)

	// 开启任务
	this.app.Action = func(ctx *cli.Context) error {
		return this.runWeb()
	}
	err := this.app.Run(os.Args)
	if err != nil {
		tog.Error(err.Error())
	}
}
