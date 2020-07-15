package event

import (
	"github.com/zhanghup/go-app/beans"
	"xorm.io/xorm"
)

const (
	xorm_default_init = "xorm_default_init"
	dict_change       = "dict:change"
	user_login        = "user:login"
	user_create       = "user:create"
	user_update       = "user:update"
	user_remove       = "user:remove"
	user_role         = "user_role:change"
)

// 数据库初始化事件
func XormDefaultInit(db *xorm.Engine)                       { EventPublish(xorm_default_init, db) }
func XormDefaultInitSubscribeOnce(fn func(db *xorm.Engine)) { EventSubscribe(xorm_default_init, fn) }

// 数据字典更新事件
func DictChangePush()               { EventPublish(dict_change) }
func DictChangeSubscribe(fn func()) { EventSubscribe(dict_change, fn) }

// 用户事件
func UserLogin(ty string, user *beans.User)            { EventPublish(user_login, ty, user) } // 登录
func UserLoginSubscribe(fn func(ty, user *beans.User)) { EventSubscribe(user_login, fn) }     // 登录【订阅】
func UserCreate(user *beans.User)                      { EventPublish(user_create, user) }    // 用户创建
func UserCreateSubscribe(fn func(user *beans.User))    { EventSubscribe(user_create, fn) }    // 用户创建【订阅】
func UserUpdate(user *beans.User)                      { EventPublish(user_update, user) }    // 用户创建
func UserUpdateSubscribe(fn func(user *beans.User))    { EventSubscribe(user_update, fn) }    // 用户创建【订阅】
func UserRemove(user *beans.User)                      { EventPublish(user_remove, user) }    // 用户创建
func UserRemoveSubscribe(fn func(user *beans.User))    { EventSubscribe(user_remove, fn) }    // 用户创建【订阅】
