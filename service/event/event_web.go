package event

import (
	"context"
	"github.com/zhanghup/go-app/beans"
)

type MsgAction string
type MsgTarget string

const (
	message          = "message"
	MsgActionAdd     = MsgAction("add")
	MsgActionRead    = MsgAction("read")
	MsgActionConfirm = MsgAction("confirm")
	MsgTargetWeb     = MsgTarget("web")
	MsgTargetApp     = MsgTarget("app")
)

type MsgInfo struct {
	Action   MsgAction
	Messages []beans.MsgInfo
}

/*
	消息事件 - 插入
*/
func MsgNew(uid string, target MsgTarget, action MsgAction, msg []beans.MsgInfo) {
	EventPublish(message+"-"+uid+"-"+string(target), MsgInfo{Action: action, Messages: msg})
}

/*
	消息事件 - 数据监听
*/
func MsgNewSubscribe(uid string, target MsgTarget, fn func(msg MsgInfo)) {
	EventSubscribe(message+"-"+uid+"-"+string(target), fn)
}

/*
	消息事件 - 取消监听
*/
func MsgNewUnSubscribe(uid string, target MsgTarget, fn func(msg MsgInfo)) {
	EventUnsubscribe(message+"-"+uid+"-"+string(target), fn)
}

/*
	消息事件 - 取消监听
*/
func MsgNewUnSubscribeWithContext(ctx context.Context, uid string, target MsgTarget, fn func(msg MsgInfo)) {
	<-ctx.Done()
	EventUnsubscribe(message+"-"+uid+"-"+string(target), fn)
}
