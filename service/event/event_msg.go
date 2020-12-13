package event

import (
	"context"
	"github.com/zhanghup/go-app/beans"
)

type MsgTarget string

const (
	message      = "message"
	MsgTargetWeb = MsgTarget("web")
	MsgTargetApp = MsgTarget("app")
)

/*
	消息事件 - 插入
*/
func MsgNew(uid string, target MsgTarget, tpl beans.MsgTemplate, msg beans.MsgInfo) {
	EventPublish(message+"-"+uid+"-"+string(target), tpl, msg)
}

/*
	消息事件 - 数据监听
*/
func MsgNewSubscribe(uid string, target MsgTarget, fn func(tpl beans.MsgTemplate, msg beans.MsgInfo)) {
	EventSubscribe(message+"-"+uid+"-"+string(target), fn)
}

/*
	消息事件 - 取消监听
*/
func MsgNewUnSubscribe(uid string, target MsgTarget, fn func(tpl beans.MsgTemplate, msg beans.MsgInfo)) {
	EventUnsubscribe(message+"-"+uid+"-"+string(target), fn)
}

/*
	消息事件 - 取消监听
*/
func MsgNewUnSubscribeWithContext(ctx context.Context, uid string, target MsgTarget, fn func(tpl beans.MsgTemplate, msg beans.MsgInfo)) {
	<-ctx.Done()
	EventUnsubscribe(message+"-"+uid+"-"+string(target), fn)
}
