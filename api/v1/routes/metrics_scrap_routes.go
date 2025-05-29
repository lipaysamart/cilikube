package routes

import (
	"github.com/ciliverse/cilikube/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterMetricsScrapeRoutes(router *gin.RouterGroup, handler *handlers.MetricsScrapeHandler) {
	// Operations within a specific namespace
	// Metrics endpoints within a namespace
	podMetricsGroup := router.Group("/namespaces/:namespace/metrics/pods")
	{
		podMetricsGroup.GET("", handler.GetPodMetricsList)   // Get Pod Metrics List
		podMetricsGroup.GET("/:name", handler.GetPodMetrics) // Get Pod Metrics
	}
	nodeMetricsGroup := router.Group("/nodes/metrics")
	{
		nodeMetricsGroup.GET("/:name", handler.GetNodeMetrics)
		nodeMetricsGroup.GET("", handler.GetNodeMetricsList)
	}
}
