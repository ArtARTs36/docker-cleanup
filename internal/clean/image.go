package clean

import (
	"context"
	"log/slog"

	"github.com/artarts36/docker-cleanup/internal/metrics"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ImageCleaner struct {
	client  *client.Client
	metrics metrics.Collector
}

func NewImageCleaner(cli *client.Client, metricsCollector metrics.Collector) *ImageCleaner {
	return &ImageCleaner{
		client:  cli,
		metrics: metricsCollector,
	}
}

func (c *ImageCleaner) Clean(ctx context.Context) error {
	report, err := c.client.ImagesPrune(ctx, filters.NewArgs())
	if err != nil {
		return err
	}

	slog.InfoContext(
		ctx,
		"[image-cleaner] cleaned images",
		slog.Any("images", report.ImagesDeleted),
		slog.Uint64("reclaimed_space", report.SpaceReclaimed),
	)

	c.metrics.ImagesCleaned(len(report.ImagesDeleted))

	return nil
}
