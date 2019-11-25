package gs

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app"
)

const (
	GIN_CONTEXT = "gin-context"
	GIN_TOKEN   = "gin-token"
)

type Middleware struct {
	gin *gin.Context
}

func NewMiddleware(g context.Context) Middleware {
	gg := g.Value(GIN_CONTEXT)
	return Middleware{gg.(*gin.Context)}
}

func (this Middleware) GinContext() *gin.Context {
	return this.gin
}

func (this Middleware) Uid() string {
	uido, ok := this.gin.Get("uid")
	if !ok {
		return ""
	}
	return uido.(string)
}

func (this Middleware) User() app.User {
	uido, ok := this.gin.Get("user")
	if !ok {
		return app.User{}
	}
	return uido.(app.User)
}

func (this Middleware) Admin() bool {
	uido, ok := this.gin.Get("admin")
	if !ok {
		return false
	}
	return uido.(bool)
}

func (this Middleware) Perms() Perms {
	uido, ok := this.gin.Get("perms")
	if !ok {
		return nil
	}
	return uido.(Perms)
}

func (this Middleware) PermObjs() PermObjects {
	uido, ok := this.gin.Get("permobjs")
	if !ok {
		return nil
	}
	return uido.(PermObjects)
}
