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
)

type Me struct {
	Gin  *gin.Context
	Info ca.User
}

func MyInfo(g context.Context) Me {
	gg := g.Value(GIN_CONTEXT)
	ggg := gg.(*gin.Context)
	user, _ := ggg.Get("user_info")
	return Me{ggg, user.(ca.User)}
}
