package ca

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"time"
	"xorm.io/xorm"
)

type User struct {
	User        *beans.User
	Token       *beans.UserToken
	PermObjects map[string]string
	Perms       map[string][]string
}

type userCache struct {
	data     tools.ICache
	tokenmap tools.ICache
	db       *xorm.Engine
}

func (this *userCache) Set(token string, user User) {
	this.data.Set(token, user, time.Now().Unix()+3600)
	this.tokenmap.Set(*user.User.Id, *user.Token.Id, time.Now().Unix()+3600)
}

func (this *userCache) Get(token string) (User, bool) {
	u := User{}
	ok := this.data.Get(token, &u)
	return u, ok
}

func (this *userCache) RemoveByToken(token string) {
	user := User{}
	ok := this.data.Get(token, &user)
	if ok {
		this.tokenmap.Delete(*user.User.Id)
	}
	this.data.Delete(token)
}
func (this *userCache) RemoveByUser(user string) {
	token := ""
	ok := this.tokenmap.Get(user, &token)
	if ok {
		this.data.Delete(token)
	}
	this.tokenmap.Delete(token)
}

var UserCache *userCache

func init() {
	event.XormDefaultInitSubscribeOnce(func(db *xorm.Engine) {
		UserCache = &userCache{
			data:     tools.CacheCreate(true),
			tokenmap: tools.CacheCreate(true),
			db:       db,
		}

		go event.UserRemoveSubscribe(func(user *beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})

		go event.UserUpdateSubscribe(func(user *beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})

		go event.UserLoginSubscribe(func(ty, user *beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})
	})
}
