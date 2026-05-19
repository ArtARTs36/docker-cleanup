package cleaner

import (
	"github.com/artarts36/docker-cleanup/internal/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type (
	MetricsCollector           metrics.Collector
	PrometheusMetricsCollector interface {
		MetricsCollector
		Describe(ch chan<- *prometheus.Desc)
		Collect(ch chan<- prometheus.Metric)
	}
)

func NewPrometheusMetricsCollector(namespace string) PrometheusMetricsCollector {
	return PrometheusMetricsCollector(metrics.NewPrometheusCollector(namespace))
}
