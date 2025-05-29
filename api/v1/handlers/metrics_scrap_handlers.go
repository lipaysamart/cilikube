package handlers

import (
	"net/http"
	"strings"

	"github.com/ciliverse/cilikube/internal/service"
	"github.com/ciliverse/cilikube/pkg/utils"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/api/errors"
)

type MetricsScrapeHandler struct {
	service *service.MetricsScrapeService
}

func NewMetricsScrapeHandler(svc *service.MetricsScrapeService) *MetricsScrapeHandler {
	return &MetricsScrapeHandler{service: svc}
}

func (h *MetricsScrapeHandler) GetPodMetrics(c *gin.Context) {
	namespace := strings.TrimSpace(c.Param("namespace"))
	podName := strings.TrimSpace(c.Param("name"))

	if !utils.ValidateNamespace(namespace) {
		respondError(c, http.StatusBadRequest, "无效的命名空间格式")
		return
	}

	podMetrics, err := h.service.GetPodMetrics(namespace, podName)
	if err != nil {
		if errors.IsNotFound(err) {
			respondError(c, http.StatusInternalServerError, "失败的获取 Metrics")
			return
		}
	}

	respondSuccess(c, http.StatusOK, podMetrics)
}

func (h *MetricsScrapeHandler) GetPodMetricsList(c *gin.Context) {
	namespace := strings.TrimSpace(c.Param("namespace"))

	if !utils.ValidateNamespace(namespace) {
		respondError(c, http.StatusBadRequest, "无效的命名空间格式")
		return
	}

	podMetricsList, err := h.service.GetPodMetricsList(namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			respondError(c, http.StatusInternalServerError, "失败的获取 MetricsList")
			return
		}
	}

	respondSuccess(c, http.StatusOK, podMetricsList)
}

func (h *MetricsScrapeHandler) GetNodeMetrics(c *gin.Context) {
	nodeName := strings.TrimSpace(c.Param("name"))
	if !utils.ValidateResourceName(nodeName) {
		respondError(c, http.StatusBadRequest, "无效的资源名称格式")
		return
	}

	nodeMetrics, err := h.service.GetNodeMetrics(nodeName)
	if err != nil {
		if errors.IsNotFound(err) {
			respondError(c, http.StatusInternalServerError, "失败的获取 Metrics")
			return
		}
	}
	respondSuccess(c, http.StatusOK, nodeMetrics)
}

func (h *MetricsScrapeHandler) GetNodeMetricsList(c *gin.Context) {
	nodeMetricsList, err := h.service.GetNodeMetricsList()
	if err != nil {
		if errors.IsNotFound(err) {
			respondError(c, http.StatusInternalServerError, "失败的获取 MetricsList")
			return
		}
	}
	respondSuccess(c, http.StatusOK, nodeMetricsList)
}
