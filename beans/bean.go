package beans

import (
	"xorm.io/xorm"
)

func Sync(db *xorm.Engine) {
	err := db.Sync2(sys_tables()...)
	if err != nil {
		panic(err)
	}
	err = db.Sync2(log_tables()...)
	if err != nil {
		panic(err)
	}

}
