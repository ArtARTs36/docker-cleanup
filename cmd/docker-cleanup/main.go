package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/artarts36/docker-cleanup/internal/clean"
	cli "github.com/artarts36/singlecli"
	"github.com/docker/docker/client"
)

var (
	Version   = "0.1.0"
	BuildDate = "2025-05-07 20:51:00"
)

func main() {
	app := &cli.App{
		BuildInfo: &cli.BuildInfo{
			Name:        "docker-cleanup",
			Description: "docker cleanup",
			Version:     Version,
			BuildDate:   BuildDate,
		},
		Opts: []*cli.OptDefinition{
			{
				Name:        "containers",
				Description: "clean containers",
			},
			{
				Name:        "images",
				Description: "clean images",
			},
		},
		Action: run,
	}

	app.RunWithGlobalArgs(context.Background())
}

func run(ctx *cli.Context) error {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("create docker client: %w", err)
	}

	cleaner := clean.CreateCleaner(dockerClient, clean.Opts{
		Containers: ctx.HasOpt("containers"),
		Images:     ctx.HasOpt("images"),
	})

	slog.InfoContext(ctx.Context, "cleaning")

	err = cleaner.Clean(ctx.Context)
	if err != nil {
		return fmt.Errorf("cleanup failed: %w", err)
	}

	return nil
}
