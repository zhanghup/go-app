package initia

import (
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

func InitUser(e *xorm.Engine) {
	ok, err := e.Table(app.User{}).Where("id = ?", "root").Exist()
	if err != nil {
		panic(err)
	}
	if ok {
		return
	}

	slat := *tools.ObjectString()
	password := tools.Password("bwg7xj98b3", slat)
	user := app.User{
		Bean: app.Bean{
			Id:     tools.Ptr().String("root"),
			Status: tools.Ptr().Int(1),
			Weight: tools.Ptr().Int(0),
		},
		Type:     tools.Ptr().String("0"), // 超级管理员
		Account:  tools.Ptr().String("root"),
		Password: &password,
		Slat:     &slat,
	}
	_, err = e.Table(user).Insert(user)
	if err != nil {
		panic(err)
	}
}
