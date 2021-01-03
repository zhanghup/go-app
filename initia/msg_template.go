package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

func InitMsgTemplate(db *xorm.Engine) {
	InitMsgTemplateCode(db, "系统消息", "system", "notice", "1", "web", "系统消息", "")
}

func InitMsgTemplateCode(db *xorm.Engine, name, code, typ, level, target, remark, templateCode string) {
	oldTpl := beans.MsgTemplate{}
	ok, err := db.Table(&oldTpl).Where("code = ?", code).Get(&oldTpl)
	if err != nil {
		tog.Error(err.Error())
		return
	}
	if !ok {
		_, err = db.Insert(beans.MsgTemplate{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.String("1"),
			},
			Name:         &name,
			Code:         &code,
			Type:         &typ,
			Level:        &level,
			Target:       &target,
			Remark:       &remark,
			Delay:        tools.Ptr.Int64(5),
			Alert:        tools.Ptr.Int64(86400), // 提前一天开始提醒
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
