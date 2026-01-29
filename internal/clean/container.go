package clean

import (
	"context"
	"log/slog"

	"github.com/artarts36/docker-cleanup/internal/metrics"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ContainerCleaner struct {
	client  *client.Client
	metrics metrics.Collector
}

func NewContainerCleaner(cli *client.Client, metricsCollector metrics.Collector) *ContainerCleaner {
	return &ContainerCleaner{
		client:  cli,
		metrics: metricsCollector,
	}
}

func (c *ContainerCleaner) Clean(ctx context.Context) error {
	report, err := c.client.ContainersPrune(ctx, filters.NewArgs())
	if err != nil {
		return err
	}

	slog.InfoContext(
		ctx,
		"[container-cleaner] cleaned containers",
		slog.Any("containers", report.ContainersDeleted),
		slog.Uint64("reclaimed_space", report.SpaceReclaimed),
	)

	c.metrics.ContainersCleaned(len(report.ContainersDeleted))

	return nil
}
