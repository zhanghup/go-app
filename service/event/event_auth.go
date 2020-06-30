package event

// ty: 登录类型 id: 用户id
func UserLogin(ty, id string) {
	EventPublish("user:login", ty, id)
}
func UserLoginSubscribe(fn func(ty, id string)) {
	EventSubscribe("user:login", fn)
}
