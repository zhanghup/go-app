package ags

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"strings"
	"time"
	"xorm.io/xorm"
)

type IMessage interface {
	Send() error
	TimeoutMark() error
}

type message struct {
	db  *xorm.Engine
	dbs *txorm.Engine
}

func NewMessage(db *xorm.Engine) IMessage {
	return &message{
		db:  db,
		dbs: txorm.NewEngine(db),
	}
}

/*
	实时消息推送

	1. 遍历当前所有未推送的数据
	2. 归类用户以及推送平台
	3. 推送数据
	4. 插入到已推送的表中
*/
func (this *message) Send() error {
	events := make([]beans.MsgEvent, 0)
	err := this.db.Table(events).Limit(500).Find(&events)
	if err != nil {
		return err
	}
	if len(events) == 0 {
		return nil
	}

	eventMap := map[string][]beans.MsgInfo{}

	nowtime := time.Now().Unix()

	for _, o := range events {
		if o.Receiver == nil || o.Target == nil {
			// TODO
			continue
		}

		for _, tag := range strings.Split(*o.Target, ",") {
			info := beans.MsgInfo{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Status: tools.Ptr.String("1"),
				},
				Event:        o.Id,
				Receiver:     o.Receiver,
				ReceiverName: o.ReceiverName,
				Template:     o.Template,
				Level:        o.Level,
				Target:       &tag,
				Timeout:      o.Timeout,
				MustConfirm:  o.MustConfirm,
				State:        tools.Ptr.String("1"), // 未读
				SendTime:     tools.Ptr.Int64(nowtime),
				Otype:        o.Otype,
				Oid:          o.Oid,
				Title:        o.Title,
				Content:      o.Content,
				ImgPath:      o.ImgPath,
			}
			eventMap[*o.Receiver+","+tag] = append(eventMap[*o.Receiver+"-"+tag], info)
		}
	}

	deleted := make([]string, 0)
	inserted := make([]interface{}, 0)
	for k, v := range eventMap {
		sends := make([]beans.MsgInfo, 0)
		for i := range v {
			deleted = append(deleted, *v[i].Event)
			// 考虑到服务器关闭之后，消息过期的情况
			if time.Now().Unix() > *v[i].Timeout {
				v[i].State = tools.Ptr.String("3") // 未读并且已经过期
			} else {
				sends = append(sends, v[i])
			}
			inserted = append(inserted, v[i])
		}

		keys := strings.Split(k, ",")
		event.MsgNew(keys[0], event.MsgTarget(keys[1]), event.MsgActionAdd, sends)
	}

	err = this.dbs.TS(func(sess *txorm.Session) error {
		err := sess.SF(`
			delete from msg_event where id in :ids
		`, map[string]interface{}{"ids": deleted}).Exec()
		if err != nil {
			return err
		}
		err = sess.Table(beans.MsgInfo{}).Insert(inserted...)
		return err
	})

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
