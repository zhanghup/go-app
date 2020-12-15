package boot

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-app/service/job"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

type Struct struct {
	name      string
	box       *rice.Box
	db        *xorm.Engine
	msg       ags.IMessage
	routerfns []func(g *gin.Engine, db *xorm.Engine)
}

/*
	box 配置文件位置
	name 服务名称
	initdb 是否初始化数据库
*/
func Boot(box *rice.Box, name string, initdb ...bool) *Struct {
	cfg.InitConfig(box)
	s := &Struct{box: box, name: name}
	if len(initdb) > 0 && !initdb[0] {
		return s
	}

	// 初始化数据库
	if s.db != nil {
		return s
	}
	e, err := txorm.NewXorm(cfg.DB)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	e.ShowSQL(true)
	s.db = e
	initia.InitDBTemplate(e)
	return s
}

// 通知数据库完成初始化 - 需要在同步完成表结构的时候推送
func (this *Struct) XormInited() *Struct {
	// 数据库初始化完成事件
	event.XormDefaultInit(this.db)
	return this
}

// 初始化数据
func (this *Struct) Init(fns ...func(db *xorm.Engine)) *Struct {
	initia.InitDict(this.db)
	initia.InitUser(this.db)
	initia.InitMsgTemplate(this.db)
	if fns != nil {
		for _, fn := range fns {
			fn(this.db)
		}
	}
	return this
}

func (this *Struct) InitTestData(fns ...func(db *xorm.Engine)) *Struct {
	initia.InitTest(this.db)
	if fns != nil {
		for _, fn := range fns {
			fn(this.db)
		}
	}
	return this
}

//  同步表结构
func (this *Struct) SyncTables(fn ...func(db *xorm.Engine)) *Struct {
	beans.Sync(this.db)
	if len(fn) > 0 {
		for _, f := range fn {
			f(this.db)
		}
	}
	return this
}

// 初始化定时任务
func (this *Struct) JobsInit() *Struct {
	err := job.InitJobs(this.db)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	return this
}

// 初始化内置的消息任务
func (this *Struct) JobsInitMessages() *Struct {
	this.msg = ags.NewMessage(this.db)
	this.Jobs("消息超时处理任务", "0 * * * * *", func(db *xorm.Engine) error {
		return this.msg.TimeoutMark()
	})
	return this
}
func (this *Struct) Jobs(name, spec string, fn func(db *xorm.Engine) error, flag ...bool) *Struct {
	err := job.AddJob(name, spec, fn, flag...)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	return this
}

// 自定义接口
func (this *Struct) Router(fn ...func(g *gin.Engine, db *xorm.Engine)) *Struct {
	this.routerfns = append(this.routerfns, fn...)
	return this
}

func (this *Struct) StartRouter() error {
	return tgin.NewGin(cfg.Web, func(g *gin.Engine) error {
		for _, fn := range this.routerfns {
			fn(g, this.db)
		}
		return nil
	})
}
