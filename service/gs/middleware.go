package gs

import (
	"context"
	"github.com/gin-gonic/gin"
)

const (
	GIN_CONTEXT = "gin-context"
	GIN_TOKEN   = "gin-token"
)

func GinContext(ctx context.Context) *gin.Context {
	c := ctx.Value(GIN_CONTEXT)
	if c != nil {
		return c.(*gin.Context)
	}
	return &gin.Context{}
}
