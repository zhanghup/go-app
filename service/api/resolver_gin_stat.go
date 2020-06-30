package api

import (
	"github.com/rcrowley/go-metrics"
)

//Report from default metric registry
func statsReport() metrics.Registry {
	return metrics.DefaultRegistry
}

//RequestStats middleware
//func StatsRequest() gin.HandlerFunc {
//	const (
//		ginLatencyMetric = "gin.latency"
//		ginStatusMetric  = "gin.status"
//		ginRequestMetric = "gin.request"
//	)
//
//	return func(c *gin.Context) {
//		start := time.Now()
//
//		c.Next()
//
//		req := metrics.GetOrRegisterMeter(ginRequestMetric, nil)
//		req.Mark(1)
//
//		latency := metrics.GetOrRegisterTimer(ginLatencyMetric, nil)
//		latency.UpdateSince(start)
//
//		status := metrics.GetOrRegisterMeter(fmt.Sprintf("%s.%d", ginStatusMetric, c.Writer.Status()), nil)
//		status.Mark(1)
//	}
//}

