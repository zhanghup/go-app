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
func MsgNew(uid, target string, msg []beans.MsgInfo) { EventPublish(msg_new+"-"+uid+"-"+target, msg) }
func MsgNewSubscribe(uid, target string, fn func(msg []beans.MsgInfo)) {
	EventSubscribe(msg_new+"-"+uid+"-"+target, fn)
}
func MsgNewUnSubscribe(uid, target string, fn func(msg []beans.MsgInfo)) {
	EventUnsubscribe(msg_new+"-"+uid+"-"+target, fn)
}

// 消息事件 - 已读
func MsgRead(uid, target string, msg []beans.MsgInfo) {
	EventPublish(msg_read+"-"+uid+"-"+target, msg)
}
func MsgReadSubscribe(uid, target string, fn func(msg []beans.MsgInfo)) {
	EventSubscribe(msg_read+"-"+uid+"-"+target, fn)
}
func MsgReadUnSubscribe(uid, target string, fn func(msg []beans.MsgInfo)) {
	EventUnsubscribe(msg_read+"-"+uid+"-"+target, fn)
}

// 消息事件 - 确认
func MsgConfirm(uid, target string, msg beans.MsgInfo) {
	EventPublish(msg_confirm+"-"+uid+"-"+target, msg)
}
func MsgConfirmSubscribe(uid, target string, fn func(msg []beans.MsgInfo)) {
	EventSubscribe(msg_confirm+"-"+uid+"-"+target, fn)
}
func MsgConfirmUnSubscribe(uid, target string, fn func(msg []beans.MsgInfo)) {
	EventUnsubscribe(msg_confirm+"-"+uid+"-"+target, fn)
}
