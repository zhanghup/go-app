package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
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

	err = txorm.NewEngine(db).TS(func(sess txorm.ISession) error {

		user := beans.User{
			Bean: beans.Bean{
				Id:     tools.Ptr.String("root"),
				Status: tools.Ptr.String("1"),
				Weight: tools.Ptr.Int(0),
			},
		}
		err := sess.Insert(user)
		if err != nil {
			return err
		}

		salt := tools.Str.Uid()
		password := tools.Crypto.Password("Aa123456.", salt)
		err = sess.Insert(beans.Account{
			Bean: beans.Bean{
				Id:     tools.Ptr.String("root"),
				Status: tools.Ptr.String("1"),
				Weight: tools.Ptr.Int(0),
			},
			Type:     tools.Ptr.String("password"),
			Uid:      user.Id,
			Username: tools.Ptr.String("root"),
			Password: &password,
			Salt:     &salt,
			Default:  tools.Ptr.Int(1),
		})
		return err
	})

	if err != nil {
		panic(err)
	}
}
