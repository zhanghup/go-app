package directive

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/ca"
)

type Me ca.User

func MyInfo(g context.Context) Me {
	gg := g.Value(GIN_CONTEXT)
	ggg := gg.(*gin.Context)
	user, _ := ggg.Get(GIN_USER)
	return Me(user.(ca.User))
}
