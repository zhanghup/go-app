package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

func InitUser(db *xorm.Engine) {
	ok, err := db.Table(beans.User{}).Where("id = ?", "root").Exist()
	if err != nil {
		panic(err)
	}
	if ok {
		return
	}

	slat := tools.Str.Uid()
	password := tools.Crypto.Password("zhang3611", slat)
	user := beans.User{
		Bean: beans.Bean{
			Id:     tools.Ptr.String("root"),
			Status: tools.Ptr.Int(1),
			Weight: tools.Ptr.Int(0),
		},
		Type:     tools.Ptr.String("0"), // 超级管理员
		Account:  tools.Ptr.String("root"),
		Password: &password,
		Salt:     &slat,
	}
	_, err = db.Table(user).Insert(user)
	if err != nil {
		panic(err)
	}
}
