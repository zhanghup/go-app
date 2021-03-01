package beans

import (
	"xorm.io/xorm"
)

type PageResult struct {
	Total int64       `json:"total"`
	Datas interface{} `json:"datas"`
}
type PageParam struct {
	Index int  `json:"index"`
	Size  int  `json:"size"`
	Count bool `json:"count"`
}

type Bean struct {
	Id      *string `json:"id" xorm:"Varchar(128) pk"`
	Created *int64  `json:"created" xorm:"created Int(14)"`
	Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
	Weight  *int    `json:"weight" xorm:"weight  Int(9) default(0)"`
	Status  *string `json:"status" xorm:"status  Int(1) default('1')"`
}

func Sync(db *xorm.Engine) {
	err := db.Sync2(sys_tables()...)
	if err != nil {
		panic(err)
	}
	err = db.Sync2(msg_tables()...)
	if err != nil {
		panic(err)
	}
	err = db.Sync2(log_tables()...)
	if err != nil {
		panic(err)
	}
	err = db.Sync2(plan_tables()...)
	if err != nil {
		panic(err)
	}
	err = db.Sync2(wxmp_tables()...)
	if err != nil {
		panic(err)
	}
}
