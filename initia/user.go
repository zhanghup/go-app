package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
)

func InitUser() {
	ok, err := gs.DB().Table(beans.User{}).Where("id = ?", "root").Exist()
	if err != nil {
		panic(err)
	}
	if ok {
		return
	}

	err = gs.DBS().TS(func(sess txorm.ISession) error {

		user := beans.User{
			Bean: beans.Bean{
				Id:     tools.PtrOfString("root"),
				Status: tools.PtrOfString("1"),
				Weight: tools.PtrOfInt(0),
			},
			Name: tools.PtrOfString("超级管理员"),
		}
		err := sess.Insert(user)
		if err != nil {
			return err
		}

		salt := tools.UUID()
		password := tools.Password("Aa123456.", salt)
		err = sess.Insert(beans.Account{
			Bean: beans.Bean{
				Id:     tools.PtrOfString("root"),
				Status: tools.PtrOfString("1"),
				Weight: tools.PtrOfInt(0),
			},
			Type:     tools.PtrOfString("password"),
			Uid:      user.Id,
			Username: tools.PtrOfString("root"),
			Password: &password,
			Salt:     &salt,
			Default:  tools.PtrOfInt(1),
		})
		return err
	})

	if err != nil {
		panic(err)
	}
}
