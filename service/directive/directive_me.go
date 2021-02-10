package directive

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/ca"
)

func MyInfo(g context.Context) *ca.User {
	gg := g.Value(GIN_CONTEXT)
	ggg := gg.(*gin.Context)
	user, _ := ggg.Get(GIN_USER)
	u := user.(ca.User)
	return &u
}

func MyWxmpUser(g context.Context) *ca.WxmpUser {
	gg := g.Value(GIN_CONTEXT)
	ggg := gg.(*gin.Context)
	user, _ := ggg.Get(GIN_WXUSER)
	u := user.(ca.WxmpUser)
	return &u
}
