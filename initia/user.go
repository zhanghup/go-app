package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-tools"
)

func initUser() {
	ok, err := cfg.DB().Engine().Table(beans.User{}).Where("id = ?", "root").Exist()
	if err != nil {
		panic(err)
	}
	if ok {
		return
	}

	slat := *tools.ObjectString()
	password := tools.Password("bwg7xj98b3", slat)
	user := beans.User{
		Bean: beans.Bean{
			Id:     tools.Ptr().String("root"),
			Status: tools.Ptr().Int(1),
			Weight: tools.Ptr().Int(0),
		},
		Type:     tools.Ptr().String("0"), // 超级管理员
		Account:  tools.Ptr().String("root"),
		Password: &password,
		Slat:     &slat,
	}
	_, err = cfg.DB().Engine().Table(user).Insert(user)
	if err != nil {
		panic(err)
	}
}
