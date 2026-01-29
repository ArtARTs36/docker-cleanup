package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusCollector struct {
	containersCleaned prometheus.Counter
	imagesCleaned     prometheus.Counter
}

func NewPrometheusCollector(namespace string) *PrometheusCollector {
	return &PrometheusCollector{
		containersCleaned: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "containers_cleaned",
		}),
		imagesCleaned: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "images_cleaned",
		}),
	}
}

func (c *PrometheusCollector) ContainersCleaned(count int) {
	c.containersCleaned.Add(float64(count))
}

func (c *PrometheusCollector) ImagesCleaned(count int) {
	c.imagesCleaned.Add(float64(count))
}

func (c *PrometheusCollector) Flush(_ context.Context) error {
	return nil
}

func (c *PrometheusCollector) Describe(ch chan<- *prometheus.Desc) {
	c.containersCleaned.Describe(ch)
	c.imagesCleaned.Describe(ch)
}

func (c *PrometheusCollector) Collect(ch chan<- prometheus.Metric) {
	c.containersCleaned.Collect(ch)
	c.imagesCleaned.Collect(ch)
}
