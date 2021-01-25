package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

func InitMsgTemplate(db *xorm.Engine) {
	InitMsgTemplateCode(db, "系统消息", "system", "notice", "1", "web", "系统消息", "", "", "1", 5, 86400)
}

func InitMsgTemplateCode(db *xorm.Engine, name, code, typ, level, target, remark, templateCode, toadmin, Template string, delay, alert int64) {
	oldTpl := beans.MsgTemplate{}
	ok, err := db.Table(&oldTpl).Where("code = ?", code).Get(&oldTpl)
	if err != nil {
		tog.Error(err.Error())
		return
	}
	if !ok {
		_, err = db.Insert(beans.MsgTemplate{
			Bean: beans.Bean{
				Id:     tools.PtrOfUUID(),
				Status: tools.PtrOfString("1"),
			},
			Name:         &name,
			Code:         &code,
			Type:         &typ,
			Level:        &level,
			Target:       &target,
			Remark:       &remark,
			Delay:        &delay,
			Alert:        &alert, // 提前一天开始提醒
			ToAdmin:      &toadmin,
			Template:     &Template,
			TemplateCode: &templateCode,
		})
		if err != nil {
			tog.Error(err.Error())
			return
		}
	} else {
		_, err = db.Table(oldTpl).Where("id = ?", oldTpl.Id).Update(map[string]interface{}{
			"template_code": templateCode,
		})
		if err != nil {
			tog.Error(err.Error())
			return
		}
	}
}
