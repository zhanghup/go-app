package msg

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"strings"
	"xorm.io/xorm"
)

type IMessage interface {
	Run() error
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

func (this *message) Run() error {
	events := make([]beans.MsgEvent, 0)
	err := this.db.Table(events).OrderBy("receiver").Find(&events)
	if err != nil {
		return err
	}

	eventMap := map[string][]beans.MsgInfo{}

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
				Event:         o.Id,
				Receiver:      o.Receiver,
				ReceiverName:  o.ReceiverName,
				Template:      o.Template,
				Level:         o.Level,
				Target:        &tag,
				Timeout:       o.Timeout,
				MustConfirm:   o.MustConfirm,
				ConfirmTarget: o.ConfirmTarget,
				State:         o.State,
				Otype:         o.Otype,
				Oid:           o.Oid,
				Title:         o.Title,
				Content:       o.Content,
				ImgPath:       o.ImgPath,
			}
			eventMap[*o.Receiver] = append(eventMap[*o.Receiver], info)
		}
	}

	return nil
}
