package ags

import (
	"github.com/zhanghup/go-app/beans"
)

// ----- 消息

// 消息发送
func MessageSend(tpl beans.MsgTemplate, uid, uname, otype, oid, defaultContent string, model map[string]string) error {
	return defaultMessage.NewMessage(tpl, uid, uname, otype, oid, defaultContent, model)
}

