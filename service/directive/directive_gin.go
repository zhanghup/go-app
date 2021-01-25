package directive

import (
	"context"
	"github.com/gin-gonic/gin"
)

func Gin(g context.Context) *gin.Context {
	gg := g.Value(GIN_CONTEXT)
	return gg.(*gin.Context)
}
