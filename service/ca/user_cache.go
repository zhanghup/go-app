package ca

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"time"
	"xorm.io/xorm"
)

type User struct {
	User        beans.User
	Account     beans.Account
	Token       beans.Token
	PermObjects map[string]string
	Admin       bool
	TokenString string
}

type userCache struct {
	data     tools.ICache
	tokenmap tools.ICache
	db       *xorm.Engine
}

func (this *userCache) Set(token string, user User) {
	this.data.Set(token, user, time.Now().Unix()+7200)
	this.tokenmap.Set(*user.User.Id, *user.Token.Id, time.Now().Unix()+7200)
}

func (this *userCache) Get(token string) (User, bool) {
	o := this.data.Get(token)
	if o == nil {
		return User{}, false
	} else {
		return o.(User), true
	}
}

func (this *userCache) GetByUser(uid string) (User, bool) {
	o := this.tokenmap.Get(uid)
	if o != nil {
		token := o.(string)
		return this.Get(token)
	}
	return User{}, false
}

func (this *userCache) RemoveByToken(token string) {
	o := this.data.Get(token)
	if o != nil {
		user := o.(User)
		this.tokenmap.Delete(*user.User.Id)
	}
	this.data.Delete(token)
}
func (this *userCache) RemoveByUser(user string) {
	o := this.tokenmap.Get(user)
	if o != nil {
		token := o.(string)
		this.data.Delete(token)
	}
	this.tokenmap.Delete(user)
}

var UserCache *userCache

func init() {
	event.XormDefaultInitSubscribeOnce(func(db *xorm.Engine) {
		UserCache = &userCache{
			data:     tools.CacheCreate(true),
			tokenmap: tools.CacheCreate(true),
			db:       db,
		}

		go event.UserRemoveSubscribe(func(user beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})

		go event.UserUpdateSubscribe(func(user beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})

		go event.UserLoginSubscribe(func(acc beans.Account, user beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})

		go event.UserRoleChangeSubscribe(func(user beans.User) {
			UserCache.RemoveByUser(*user.Id)
		})
	})
}
