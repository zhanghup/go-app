package event

import "github.com/zhanghup/go-app/beans"

const (
	user_login = "user:login"
)

// ty: 登录类型 id: 用户id
func UserLogin(ty string, user beans.User) {
	EventPublish(user_login, ty, user)
}
func UserLoginSubscribe(fn func(ty, user beans.User)) {
	EventSubscribe(user_login, fn)
}
