package directive

import (
	"context"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/ca"
)

func MyInfo(g context.Context) *ca.User {
	user, _ := gs.Gin(g).Get(gs.GIN_USER)
	u := user.(ca.User)
	return &u
}

func MyWxmpUser(g context.Context) *ca.WxmpUser {
	user, _ := gs.Gin(g).Get(gs.GIN_WXUSER)
	u := user.(ca.WxmpUser)
	return &u
}
