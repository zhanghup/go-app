package event

import (
	"github.com/zhanghup/go-app/beans"
)

const (
	msg_new     = "msg:new"
	msg_read    = "msg:read"
	msg_confirm = "msg:confirm"
)

// 消息事件 - 新增
func MsgNew(uid string, msg beans.Msg)                     { EventPublish(msg_new+"-"+uid, msg) }
func MsgNewSubscribe(uid string, fn func(msg beans.Msg))   { EventSubscribe(msg_new+"-"+uid, fn) }
func MsgNewUnSubscribe(uid string, fn func(msg beans.Msg)) { EventUnsubscribe(msg_new+"-"+uid, fn) }

// 消息事件 - 已读
func MsgRead(uid string, msg beans.Msg)                     { EventPublish(msg_read+"-"+uid, msg) }
func MsgReadSubscribe(uid string, fn func(msg beans.Msg))   { EventSubscribe(msg_read+"-"+uid, fn) }
func MsgReadUnSubscribe(uid string, fn func(msg beans.Msg)) { EventUnsubscribe(msg_read+"-"+uid, fn) }

// 消息事件 - 确认
func MsgConfirm(uid string, msg beans.Msg)                   { EventPublish(msg_confirm+"-"+uid, msg) }
func MsgConfirmSubscribe(uid string, fn func(msg beans.Msg)) { EventSubscribe(msg_confirm+"-"+uid, fn) }
