package directive

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/ca"
)

const (
	GIN_CONTEXT       = "gin-context"
	GIN_TOKEN         = "gin-token"
	GIN_AUTHORIZATION = "Authorization"
	GIN_USER          = "user_info"
)

type Me struct {
	Gin  *gin.Context
	Info ca.User
}

func MyInfo(g context.Context) Me {
	gg := g.Value(GIN_CONTEXT)
	ggg := gg.(*gin.Context)
	user, _ := ggg.Get(GIN_USER)
	return Me{ggg, user.(ca.User)}
}
