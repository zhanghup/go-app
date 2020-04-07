package directive

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
)

const (
	GIN_CONTEXT = "gin-context"
	GIN_TOKEN   = "gin-token"
)

type Me struct {
	gin *gin.Context
}

func MyInfo(g context.Context) Me {
	gg := g.Value(GIN_CONTEXT)
	return Me{gg.(*gin.Context)}
}

func (this Me) GinContext() *gin.Context {
	return this.gin
}

func (this Me) Uid() string {
	uido, ok := this.gin.Get("uid")
	if !ok {
		return ""
	}
	return uido.(string)
}

func (this Me) User() beans.User {
	uido, ok := this.gin.Get("user")
	if !ok {
		return beans.User{}
	}
	return uido.(beans.User)
}

func (this Me) Admin() bool {
	uido, ok := this.gin.Get("admin")
	if !ok {
		return false
	}
	return uido.(bool)
}

func (this Me) Perms() Perms {
	uido, ok := this.gin.Get("perms")
	if !ok {
		return nil
	}
	return uido.(Perms)
}

func (this Me) PermObjs() PermObjects {
	uido, ok := this.gin.Get("permobjs")
	if !ok {
		return nil
	}
	return uido.(PermObjects)
}
