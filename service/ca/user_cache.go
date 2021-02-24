package ca

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"strings"
	"time"
)

type User struct {
	Id          string
	Name        string
	Admin       bool
	TokenString string

	User        beans.User
	Account     beans.Account
	Token       beans.Token
	permObjects map[string]map[string]bool
}

func (this *User) PermObjects() map[string]map[string]bool {
	return this.permObjects
}

func (this *User) EntityPermAdd(entity, mask string) {

	if this.permObjects == nil {
		this.permObjects = map[string]map[string]bool{}
	}

	if _, ok := this.permObjects[entity]; !ok {
		this.permObjects[entity] = map[string]bool{}
	}

	if mask == "" {
		return
	}

	ms := strings.Split(mask, ",")
	for _, s := range ms {
		this.permObjects[entity][s] = true
	}
}

func (this *User) EntityPerm(entity, opt string) bool {
	ent, ok := this.permObjects[entity]
	if !ok {
		return true
	}
	o, ok := ent[opt]
	return o && ok
}

type userCache struct {
	usermap  tools.ICache
	tokenmap tools.ICache
}

func (this *userCache) Set(token string, user User) {
	this.usermap.Set(token, user, time.Now().Unix()+7200)
	this.tokenmap.Set(*user.User.Id, *user.Token.Id, time.Now().Unix()+7200)
}

func (this *userCache) Get(token string) (User, bool) {
	o := this.usermap.Get(token)
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
	o := this.usermap.Get(token)
	if o != nil {
		user := o.(User)
		this.tokenmap.Delete(*user.User.Id)
	}
	this.usermap.Delete(token)
}
func (this *userCache) RemoveByUser(user string) {
	o := this.tokenmap.Get(user)
	if o != nil {
		token := o.(string)
		this.usermap.Delete(token)
	}
	this.tokenmap.Delete(user)
}

func (this *userCache) Clear() {
	this.usermap.Clear()
	this.tokenmap.Clear()
}

var UserCache *userCache

func initWebUser() {
	gs.InfoBegin("web用户缓存")
	UserCache = &userCache{
		usermap:  tools.CacheCreate(true),
		tokenmap: tools.CacheCreate(true),
	}
	gs.InfoSuccess("web用户缓存")

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
}
