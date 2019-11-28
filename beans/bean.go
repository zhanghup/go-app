package beans

import (
	"github.com/zhanghup/go-app/cfg"
)

func Sync() {
	err := cfg.DB().Engine().Sync2(sys_tables()...)
	if err != nil {
		panic(err)
	}
	err = cfg.DB().Engine().Sync2(log_tables()...)
	if err != nil {
		panic(err)
	}
}
