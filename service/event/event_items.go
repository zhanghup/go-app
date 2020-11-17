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
	user_role         = "user:role"
)

// 数据库初始化事件
func XormDefaultInit(db *xorm.Engine)                       { EventPublish(xorm_default_init, db) }
func XormDefaultInitSubscribeOnce(fn func(db *xorm.Engine)) { EventSubscribe(xorm_default_init, fn) }

// 数据字典更新事件
func DictChange()                   { EventPublish(dict_change) }
func DictChangeSubscribe(fn func()) { EventSubscribe(dict_change, fn) }

// 用户事件
func UserLogin(acc beans.Account, user beans.User)                   { EventPublish(user_login, acc, user) } // 登录
func UserLoginSubscribe(fn func(acc beans.Account, user beans.User)) { EventSubscribe(user_login, fn) }      // 登录【订阅】
func UserCreate(user beans.User)                                    { EventPublish(user_create, user) }     // 用户创建
func UserCreateSubscribe(fn func(user beans.User))                  { EventSubscribe(user_create, fn) }     // 用户创建【订阅】
func UserUpdate(user beans.User)                                    { EventPublish(user_update, user) }     // 用户更新
func UserUpdateSubscribe(fn func(user beans.User))                  { EventSubscribe(user_update, fn) }     // 用户更新【订阅】
func UserRemove(user beans.User)                                    { EventPublish(user_remove, user) }     // 用户删除
func UserRemoveSubscribe(fn func(user beans.User))                  { EventSubscribe(user_remove, fn) }     // 用户删除【订阅】
func UserRoleChange(user beans.User)                                { EventPublish(user_role, user) }       // 用户角色更新
func UserRoleChangeSubscribe(fn func(user beans.User))              { EventSubscribe(user_role, fn) }       // 用户角色更新【订阅】
