package ca

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"time"
	"xorm.io/xorm"
)

type WxmpUser struct {
	Id          string
	Mobile      string
	Appid       string
	Unionid     string
	Openid      string
	Nickname    string
	TokenString string

	SessionKey string

	User *beans.WxmpUser
}

type wxmpUserCache struct {
	usermap tools.ICache
	db      *xorm.Engine
}

func (this *wxmpUserCache) Set(uid string, user WxmpUser) {
	this.usermap.Set(uid, user, time.Now().Unix()+7200)
}

func (this *wxmpUserCache) Get(uid string) (WxmpUser, bool) {
	o := this.usermap.Get(uid)
	if o == nil {
		return WxmpUser{}, false
	} else {
		return o.(WxmpUser), true
	}
}

func (this *wxmpUserCache) Remove(uid string) {
	this.usermap.Delete(uid)
}

func (this *wxmpUserCache) Clear() {
	this.usermap.Clear()
}

var WxuserCache *wxmpUserCache

func initWxmpuser() {
	gs.InfoBegin("微信小程序用户缓存")
	WxuserCache = &wxmpUserCache{
		usermap: tools.CacheCreate(true),
	}
	gs.InfoSuccess("微信小程序用户缓存")

	go event.WxmpUserRemoveSubscribe(func(user beans.WxmpUser) {
		WxuserCache.Remove(*user.Id)
	})

	go event.WxmpUserUpdateSubscribe(func(user beans.WxmpUser) {
		WxuserCache.Remove(*user.Id)
	})

	go event.WxmpUserLoginSubscribe(func(user beans.WxmpUser) {
		WxuserCache.Remove(*user.Id)
	})
}
