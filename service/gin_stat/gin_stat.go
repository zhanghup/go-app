package gin_stat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/directive"
	"net/http"
	"time"
)

//Report from default metric registry
func statsReport() metrics.Registry {
	return metrics.DefaultRegistry
}

//RequestStats middleware
func statsRequest() gin.HandlerFunc {
	const (
		ginLatencyMetric = "gin.latency"
		ginStatusMetric  = "gin.status"
		ginRequestMetric = "gin.request"
	)

	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		req := metrics.GetOrRegisterMeter(ginRequestMetric, nil)
		req.Mark(1)

		latency := metrics.GetOrRegisterTimer(ginLatencyMetric, nil)
		latency.UpdateSince(start)

		status := metrics.GetOrRegisterMeter(fmt.Sprintf("%s.%d", ginStatusMetric, c.Writer.Status()), nil)
		status.Mark(1)
	}
}

func Gin() {
	cfg.Web().Engine().Use(statsRequest())
	cfg.Web().Engine().Group("/", directive.UserAuth()).GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, statsReport())
	})
}
