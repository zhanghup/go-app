package ags

import (
	"errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"strings"
	"time"
	"xorm.io/xorm"
)

type IMessage interface {
	NewMessage(tpl beans.MsgTemplate, uid, uname, otype, oid, title, content string) error
	TimeoutMark() error
}

type message struct {
	db  *xorm.Engine
	dbs *txorm.Engine
}

/*
	实时消息推送
*/
func (this *message) NewMessage(tpl beans.MsgTemplate, uid, uname, otype, oid, title, content string) error {

	if tpl.Target == nil || len(*tpl.Target) == 0 {
		return errors.New("消息未指定需要推送的平台")
	}

	tags := strings.Split(*tpl.Target, ",")

	infos := make([]interface{}, 0)

	for j := range tags {

		nowtime := time.Now().Unix()
		timeout := int64(86400 * 365 * 100) // 一百年过期
		if tpl.Expire != nil {
			timeout = nowtime + *tpl.Expire
		}

		info := beans.MsgInfo{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.String("1"),
			},
			Receiver:     &uid,
			ReceiverName: &uname,
			Type:         tpl.Type,
			Template:     tpl.Id,
			Level:        tpl.Level,
			Target:       &tags[j],
			Timeout:      &timeout,
			MustConfirm:  tpl.MustConfirm,
			State:        tools.Ptr.String("1"), // 未读
			SendTime:     tools.Ptr.Int64(nowtime),
			Otype:        &otype,
			Oid:          &oid,
			Title:        &title,
			Content:      &content,
			ImgPath:      tpl.ImgPath,
			Remark:       tpl.Remark,
		}

		event.MsgNew(uid, event.MsgTarget(tags[j]), event.MsgActionAdd, info)
		infos = append(infos, info)
	}

	_, err := this.db.Insert(infos...)
	return err
}

/*
	历史消息标记
	将历史的已过期的数据标记为已过期状态，包含已读已过期和未读已过期
*/
func (this *message) TimeoutMark() error {
	return this.dbs.SF(`
		update 
			msg_info mi
		set state = case 
			when mi.state = '0' then '2'  
			when mi.state = '1' then '3'
		end
		where mi.timeout < unix_timestamp(now())
	`).Exec()
}
