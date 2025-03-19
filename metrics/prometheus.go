package metrics

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// 请求计数器
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// 请求处理时间
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 3, 5, 7, 10},
		},
		[]string{"method", "path"},
	)

	// 活跃请求数
	httpRequestsInProgress = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Number of HTTP requests in progress",
		},
		[]string{"method", "path"},
	)
)

// PrometheusMiddleware Gin 中间件，用于收集 Prometheus 指标
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if path == "" {
			path = "undefined"
		}

		method := c.Request.Method

		// 记录开始时间
		startTime := time.Now()

		// 增加进行中的请求计数
		httpRequestsInProgress.WithLabelValues(method, path).Inc()

		// 处理请求
		c.Next()

		// 减少进行中的请求计数
		httpRequestsInProgress.WithLabelValues(method, path).Dec()

		// 记录请求持续时间
		duration := time.Since(startTime).Seconds()
		httpRequestDuration.WithLabelValues(method, path).Observe(duration)

		// 记录请求状态
		status := string(rune(c.Writer.Status()))
		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
	}
}
