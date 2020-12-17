package ags

import (
	"github.com/zhanghup/go-app/service/event"
	"xorm.io/xorm"
)

var defaultDB *xorm.Engine

/*
	设置默认数据库引擎
*/
func DefaultDBSet(db *xorm.Engine) {
	defaultDB = db
}

func init() {
	event.XormDefaultInitSubscribeOnce(func(db *xorm.Engine) {
		defaultDB = db
		MessageInit()
		UploaderNew()
	})
}
