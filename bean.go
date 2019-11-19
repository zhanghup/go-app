package app

import "github.com/go-xorm/xorm"

func Sync(e *xorm.Engine) {
	err := e.Sync2(sys_tables()...)
	if err != nil {
		panic(err)
	}
	err = e.Sync2(log_tables()...)
	if err != nil {
		panic(err)
	}
}
