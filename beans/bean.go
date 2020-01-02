package beans

import (
	"github.com/zhanghup/go-app/ctx"
)

func Sync() {
	err := ctx.DB().Engine().Sync2(sys_tables()...)
	if err != nil {
		panic(err)
	}
	err = ctx.DB().Engine().Sync2(log_tables()...)
	if err != nil {
		panic(err)
	}

	if ctx.WxmiEnable() {
		err = ctx.DB().Engine().Sync2(wxmi_tables()...)
		if err != nil {
			panic(err)
		}
	}
}
