package clean

import (
	"context"
	"log/slog"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ImageCleaner struct {
	client *client.Client
}

func NewImageCleaner(cli *client.Client) *ImageCleaner {
	return &ImageCleaner{
		client: cli,
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

	return nil
}
