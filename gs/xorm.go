package gs

import (
	"context"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

var defaultDB *xorm.Engine
var defaultDBA txorm.IEngine

func DB() *xorm.Engine {
	if defaultDB == nil {
		panic("【数据库[DB]】未初始化！！！")
	}
	return defaultDB
}

func DBA() txorm.IEngine {
	if defaultDBA == nil {
		panic("【数据库[DBA]】未初始化！！！")
	}
	return defaultDBA
}

func DBS(ctx context.Context) txorm.ISession {
	if defaultDBA == nil {
		panic("【数据库[DBS]】未初始化！！！")
	}
	return defaultDBA.NewSession(true, ctx)
}

func Sess(ctx context.Context) txorm.ISession {
	sess := defaultDBA.Session(ctx)
	err := sess.Begin()
	if err != nil {
		tog.Error("【数据库[SESSION]】开启事务异常！！！")
	}
	return sess
}
