package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/giter/go.rice"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-app/service/file"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

type Struct struct {
	box       *rice.Box
	db        *xorm.Engine
	routerfns []func(g *gin.Engine)
}

var defaultStruct *Struct

func DefaultStruct() *Struct {
	if defaultStruct != nil {
		return defaultStruct
	}
	panic("boot 未初始化完成")
}

func Boot(box *rice.Box) *Struct {
	cfg.InitConfig(box)
	s := Struct{box: box}
	if defaultStruct == nil {
		defaultStruct = &s
	}
	return &s
}

// 初始化数据库
func (this *Struct) EnableXorm() *Struct {
	e, err := txorm.NewXorm(cfg.DB)
	e.ShowSQL(true)
	if err != nil {
		tog.Error(err.Error())
		panic(err)
	}
	this.db = e

	// 数据库初始化完成事件
	event.XormDefaultInit(e)
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
	initia.InitAction(this.db)
	if len(fn) > 0 {
		for _, f := range fn {
			f(this.db)
		}
	}
	return this
}

// 文件操作接口
func (this *Struct) RouterFile() *Struct {
	this.routerfns = append(this.routerfns, func(g *gin.Engine) {
		file.Gin(g.Group("/"), g.Group("/"), this.db)
	})
	return this
}

// 登录登出等接口
func (this *Struct) RouterAuth() *Struct {
	this.routerfns = append(this.routerfns, func(g *gin.Engine) {
		auth.Gin(g.Group("/"), this.db)
	})
	return this
}

// 内置api接口
func (this *Struct) RouterApi() *Struct {
	this.routerfns = append(this.routerfns, func(g *gin.Engine) {
		api.Gin(g.Group("/"), this.db)
	})
	return this
}

// 自定义接口
func (this *Struct) RouterOther(fn ...func(g *gin.Engine)) *Struct {
	this.routerfns = append(this.routerfns, fn...)
	return this
}

func (this *Struct) StartRouter() error {
	return tgin.NewGin(cfg.Web, func(g *gin.Engine) error {
		for _, fn := range this.routerfns {
			fn(g)
		}
		return nil
	})
}
