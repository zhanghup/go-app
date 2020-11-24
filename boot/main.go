package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/giter/go.rice"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-app/service/job"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

type Struct struct {
	box       *rice.Box
	db        *xorm.Engine
	msg       ags.IMessage
	routerfns []func(g *gin.Engine, db *xorm.Engine)
}

func Boot(box *rice.Box, initdb ...bool) *Struct {
	cfg.InitConfig(box)
	s := &Struct{box: box}
	if len(initdb) > 0 && !initdb[0] {
		return s
	}
	return s.enableXorm()
}

// 初始化数据库
func (this *Struct) enableXorm() *Struct {
	if this.db != nil {
		return this
	}
	e, err := txorm.NewXorm(cfg.DB)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	e.ShowSQL(true)
	this.db = e
	return this
}

// 通知数据库完成初始化
func (this *Struct) XormInited() *Struct {
	// 数据库初始化完成事件
	event.XormDefaultInit(this.db)
	return this
}

func (this *Struct) Init(fn func(db *xorm.Engine)) *Struct {
	fn(this.db)
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

// 初始化数据
func (this *Struct) InitDatas(fn ...func(db *xorm.Engine)) *Struct {
	initia.InitDict(this.db)
	initia.InitUser(this.db)
	if len(fn) > 0 {
		for _, f := range fn {
			f(this.db)
		}
	}
	return this
}
func (this *Struct) InitDict(ty string, dicts []initia.DictInfo) *Struct {
	initia.InitDictCode(this.db, ty, dicts)
	return this
}

// 基础操作接口
func (this *Struct) RouterAgs() *Struct {
	this.routerfns = append(this.routerfns, func(g *gin.Engine, db *xorm.Engine) {
		ags.Gin(g.Group("/"), g.Group("/"), db)
	})
	return this
}

// 内置api接口
func (this *Struct) RouterApi() *Struct {
	this.routerfns = append(this.routerfns, func(g *gin.Engine, db *xorm.Engine) {
		api.Gin(g.Group("/"), this.db)
	})
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
	this.Jobs("消息实时推送任务", "* * * * * *", this.msg.Send)
	this.Jobs("消息超时处理任务", "* * * * * *", this.msg.TimeoutMark)
	return this
}
func (this *Struct) Jobs(name, spec string, fn func() error, flag ...bool) *Struct {
	err := job.AddJob(name, spec, fn, flag...)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	return this
}

// 自定义接口
func (this *Struct) RouterOther(fn ...func(g *gin.Engine, db *xorm.Engine)) *Struct {
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
