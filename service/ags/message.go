package ags

import (
	"errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"strings"
	"time"
	"xorm.io/xorm"
)

type IMessage interface {
	NewMessage(tpl beans.MsgTemplate, uid, uname, otype, oid, defaultContent string, model map[string]string) error
}

type message struct {
	db  *xorm.Engine
	dbs txorm.IEngine
}

/*
	实时消息推送
*/
func (this *message) NewMessage(tpl beans.MsgTemplate, uid, uname, otype, oid, defaultContent string, model map[string]string) error {

	if tpl.Target == nil || len(*tpl.Target) == 0 {
		return errors.New("消息未指定需要推送的平台")
	}

	tags := strings.Split(*tpl.Target, ",")

	nowtime := time.Now().Unix()

	content := defaultContent
	if tpl.Template != nil {
		content = tools.StrTmp(*tpl.Template, model).String()
	}

	title := tools.StrTmp(*tpl.Name, model).String()

	info := beans.MsgInfo{
		Bean: beans.Bean{
			Id:     tools.PtrOfUUID(),
			Status: tools.PtrOfString("1"),
		},
		Receiver:     &uid,
		ReceiverName: &uname,
		Type:         tpl.Type,
		Template:     tpl.Id,
		Level:        tpl.Level,
		Target:       tpl.Target,
		State:        tools.PtrOfString("1"), // 未读
		SendTime:     tools.PtrOfInt64(nowtime),
		Otype:        &otype,
		Oid:          &oid,
		Title:        &title,
		Content:      &content,
		Model:        tools.PtrOfString(tools.JSONString(model)),
		ImgPath:      tpl.ImgPath,
		Remark:       tpl.Remark,
	}

	// 找到当前是否有已经存在的消息
	oldInfo := beans.MsgInfo{}
	ok, err := this.db.Where("receiver = ? and otype = ? and oid = ?", uid, otype, oid).Get(&oldInfo)
	if err != nil {
		tog.Error("【消息推送】 Error: " + err.Error())
		return err
	}

	// 更新消息或者插入消息
	if ok {
		if oldInfo.State == nil {
			oldInfo.State = tools.PtrOfString("1")
		}
		/*
			已读的消息分2中情况
				a)	（消息、通知） 将不再推送
				b)	（确认框） 将继续推送
		*/
		if *oldInfo.State == "0" && tools.StrContains([]string{"message", "notice"}, *tpl.Type) {
			return nil
		}
		// 未确认的消息将一直推送，直到确认为止
		if *oldInfo.State == "4" && tools.StrContains([]string{"confirm"}, *tpl.Type) {
			return nil
		}

		info.Id = oldInfo.Id
		if *oldInfo.State == "0" {
			info.State = oldInfo.State
		}
		_, err = this.db.Where("id = ?", info.Id).Update(info)
		if err != nil {
			tog.Error("【消息推送】 Error: " + err.Error())
			return err
		}
	} else {
		info.State = tools.PtrOfString("1") // 未读
		_, err := this.db.Insert(info)
		if err != nil {
			tog.Error("【消息推送】 Error: " + err.Error())
			return err
		}
	}

	// 新增一条历史记录
	history := beans.MsgHistory{
		Bean: beans.Bean{
			Id:     tools.PtrOfUUID(),
			Status: tools.PtrOfString("1"),
		},
		Info:         info.Id,
		Receiver:     info.Receiver,
		ReceiverName: info.ReceiverName,
		Type:         info.Type,
		Template:     info.Template,
		Level:        info.Level,
		Target:       info.Target,
		State:        info.State, // 未读
		SendTime:     info.SendTime,
		Otype:        info.Otype,
		Oid:          info.Oid,
		Title:        info.Title,
		Content:      info.Content,
		Model:        info.Model,
		ImgPath:      info.ImgPath,
		Remark:       info.Remark,
	}
	_, err = this.db.Insert(history)
	if err != nil {
		tog.Error("【消息推送】 Error: " + err.Error())
		return err
	}

	for j := range tags {
		event.MsgNew(uid, event.MsgTarget(tags[j]), tpl, info)
	}

	return nil
}

var defaultMessage IMessage

/*
	初始化消息工具
	@db: db为空，初始化默认消息工具
		 db不为空，返回一个新的消息工具，但是默认的不会被替换
*/
func MessageInit() IMessage {
	if defaultMessage != nil {
		return defaultMessage
	}

	defaultMessage = &message{
		db:  defaultDB,
		dbs: txorm.NewEngine(defaultDB),
	}

	return defaultMessage

}
