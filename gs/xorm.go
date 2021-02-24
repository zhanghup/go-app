package gs

import (
	"context"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

var defaultDB *xorm.Engine
var defaultDBS txorm.IEngine

func DB() *xorm.Engine {
	if defaultDB == nil {
		panic("【数据库[DB]】未初始化！！！")
	}
	return defaultDB
}

func DBS() txorm.IEngine {
	if defaultDBS == nil {
		panic("【数据库[DBS]】未初始化！！！")
	}
	return defaultDBS
}

func Sess(ctx context.Context) txorm.ISession {
	sess := defaultDBS.Session(ctx)
	err := sess.Begin()
	if err != nil {
		tog.Error("【数据库[DBS]】开启事务异常！！！")
	}
	return sess
}
