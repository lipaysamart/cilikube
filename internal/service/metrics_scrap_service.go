package service

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type MetricsScrapeService struct {
	mclient *versioned.Clientset
	config  *rest.Config
}

func NewMetricsScrapeService(config *rest.Config) *MetricsScrapeService {
	mclient, _ := versioned.NewForConfig(config)

	return &MetricsScrapeService{
		mclient: mclient,
		config:  config,
	}
}

// GetPodMetric 获取 Pod 的指标
func (s *MetricsScrapeService) GetPodMetrics(namespace string, podName string) (*v1beta1.PodMetrics, error) {
	podMetrics, err := s.mclient.MetricsV1beta1().PodMetricses(namespace).Get(
		context.TODO(),
		podName,
		metav1.GetOptions{},
	)

	if err != nil {
		return nil, err
	}

	return podMetrics, nil
}

func (s *MetricsScrapeService) GetPodMetricsList(namespace string) (*v1beta1.PodMetricsList, error) {
	if namespace == "all" {
		namespace = metav1.NamespaceAll
		podMetricsList, err := s.mclient.MetricsV1beta1().PodMetricses(namespace).List(
			context.TODO(),
			metav1.ListOptions{},
		)

		if err != nil {
			return nil, err
		}

		return podMetricsList, nil
	}

	podMetricsList, err := s.mclient.MetricsV1beta1().PodMetricses(namespace).List(
		context.TODO(),
		metav1.ListOptions{},
	)

	if err != nil {
		return nil, err
	}

	return podMetricsList, nil
}

func (s *MetricsScrapeService) GetNodeMetricsList() (*v1beta1.NodeMetricsList, error) {
	nodeMetricsList, err := s.mclient.MetricsV1beta1().NodeMetricses().List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return nodeMetricsList, nil
}

func (s *MetricsScrapeService) GetNodeMetrics(nodeName string) (*v1beta1.NodeMetrics, error) {
	nodeMetrics, err := s.mclient.MetricsV1beta1().NodeMetricses().Get(
		context.TODO(),
		nodeName,
		metav1.GetOptions{},
	)
	if err != nil {
		return nil, err
	}
	return nodeMetrics, nil
}
