package event

import (
	"github.com/zhanghup/go-app/beans"
)

const (
	wxmp_login  = "wxmp:login"
	wxmp_create = "wxmp:create"
	wxmp_update = "wxmp:update"
	wxmp_remove = "wxmp:remove"
	wxmp_role   = "wxmp:role"
	wxmp_pay    = "wxmp:pay_callback"
)

// 微信用户事件
func WxmpUserLogin(wxmp beans.WxmpUser)                    { EventPublish(wxmp_login, wxmp) }      // 登录
func WxmpUserLoginSubscribe(fn func(wxmp beans.WxmpUser))  { EventAsycSubscribe(wxmp_login, fn) }  // 登录【订阅】
func WxmpUserCreate(wxmp beans.WxmpUser)                   { EventPublish(wxmp_create, wxmp) }     // 用户创建
func WxmpUserCreateSubscribe(fn func(wxmp beans.WxmpUser)) { EventAsycSubscribe(wxmp_create, fn) } // 用户创建【订阅】
func WxmpUserUpdate(wxmp beans.WxmpUser)                   { EventPublish(wxmp_update, wxmp) }     // 用户更新
func WxmpUserUpdateSubscribe(fn func(wxmp beans.WxmpUser)) { EventAsycSubscribe(wxmp_update, fn) } // 用户更新【订阅】
func WxmpUserRemove(wxmp beans.WxmpUser)                   { EventPublish(wxmp_remove, wxmp) }     // 用户删除
func WxmpUserRemoveSubscribe(fn func(wxmp beans.WxmpUser)) { EventAsycSubscribe(wxmp_remove, fn) } // 用户删除【订阅】

// 微信小程序支付推送
func WxmpPayCallbackPush(order beans.WxmpOrder)               { EventPublish(wxmp_pay, order) }    // 用户支付
func WxmpPayCallbackSubscribe(fn func(order beans.WxmpOrder)) { EventAsycSubscribe(wxmp_pay, fn) } // 用户支付【订阅】
