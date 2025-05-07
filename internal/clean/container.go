package clean

import (
	"context"
	"log/slog"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ContainerCleaner struct {
	client *client.Client
}

func NewContainerCleaner(cli *client.Client) *ContainerCleaner {
	return &ContainerCleaner{
		client: cli,
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

	return nil
}
