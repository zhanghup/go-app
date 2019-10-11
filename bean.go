package app

import "github.com/go-xorm/xorm"

// 菜单显示
type Menu struct {
	Bean `xorm:"extends"`

	Title  *string `json:"title"`
	Name   *string `json:"name"`
	NameEn *string `json:"name_en"`
	Index  *string `json:"index"`
	Icon   *string `json:"icon"`
	Parent *string `json:"parent"`
}

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
