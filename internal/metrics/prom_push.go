package metrics

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type pushPrometheusCollector struct {
	collector Collector

	pusher *push.Pusher
}

func NewPushPrometheusCollector(
	collector Collector,
	url string,
	gatherer prometheus.Gatherer,
) Collector {
	return &pushPrometheusCollector{
		collector: collector,
		pusher:    push.New(url, "dockercleanup").Gatherer(gatherer),
	}
}

func (c *pushPrometheusCollector) ContainersCleaned(count int) {
	c.collector.ContainersCleaned(count)
}

func (c *pushPrometheusCollector) ImagesCleaned(count int) {
	c.collector.ImagesCleaned(count)
}

func (c *pushPrometheusCollector) Flush(ctx context.Context) error {
	err := c.pusher.PushContext(ctx)
	if err != nil {
		return fmt.Errorf("push metrics to push-gateway: %w", err)
	}

	slog.InfoContext(ctx, "metrics pushed to push-gateway")

	err = c.collector.Flush(ctx)
	if err != nil {
		return fmt.Errorf("flush decorated collector: %w", err)
	}

	return nil
}
