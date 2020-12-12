package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

func InitMsgTemplate(db *xorm.Engine) {
	InitMsgTemplateCode(db, "system")
}

func InitMsgTemplateCode(db *xorm.Engine, codes ...string) {
	for _, code := range codes {
		ok, err := db.Table(beans.MsgTemplate{}).Where("code = ?", code).Exist()
		if err != nil {
			tog.Error(err.Error())
			continue
		}
		if !ok {
			_, err = db.Insert(beans.MsgTemplate{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Status: tools.Ptr.String("1"),
				},
				Name:        tools.Ptr.String("系统消息"),
				Code:        &code,
				Type:        tools.Ptr.String("message"),
				Level:       tools.Ptr.String("3"),
				Target:      tools.Ptr.String("web"),
				Expire:      nil,
				MustConfirm: tools.Ptr.String("0"),
				Remark:      tools.Ptr.String("系统消息"),
			})
			if err != nil {
				tog.Error(err.Error())
				continue
			}
		}

	}
}
