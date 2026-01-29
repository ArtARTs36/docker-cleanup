package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/artarts36/docker-cleanup/internal/clean"
	"github.com/artarts36/docker-cleanup/internal/metrics"
	cli "github.com/artarts36/singlecli"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
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
			{
				Name:        "metrics-server",
				Description: "collecting metrics when `metrics-server` option filled",
				WithValue:   true,
			},
		},
		Action: run,
	}

	app.RunWithGlobalArgs(context.Background())
}

func run(cliCtx *cli.Context) error {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("create docker client: %w", err)
	}
	defer func() {
		err = dockerClient.Close()
		if err != nil {
			slog.ErrorContext(cliCtx.Context, "failed to close docker client", slog.Any("err", err))
		}
	}()

	metricsCollector, err := createMetricsCollector(cliCtx.Opts["metrics-server"])
	if err != nil {
		return fmt.Errorf("create metrics collector: %w", err)
	}

	cleaner := clean.CreateCleaner(dockerClient, clean.Opts{
		Containers: cliCtx.HasOpt("containers"),
		Images:     cliCtx.HasOpt("images"),
	}, metricsCollector)

	slog.InfoContext(cliCtx.Context, "cleaning")

	err = cleaner.Clean(cliCtx.Context)
	if err != nil {
		return fmt.Errorf("cleanup failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(cliCtx.Context, time.Minute)
	defer cancel()

	err = metricsCollector.Flush(ctx)
	if err != nil {
		return fmt.Errorf("flush metrics: %w", err)
	}

	return nil
}

func createMetricsCollector(metricsServer string) (metrics.Collector, error) {
	if metricsServer == "" {
		return metrics.NoopCollector{}, nil
	}

	registry := prometheus.NewRegistry()
	collector := metrics.NewPrometheusCollector("dockercleanup")

	if err := registry.Register(collectors.NewBuildInfoCollector()); err != nil {
		return nil, fmt.Errorf("register build info: %w", err)
	}

	if err := registry.Register(collector); err != nil {
		return nil, fmt.Errorf("register prometheus collector: %w", err)
	}

	return metrics.NewPushPrometheusCollector(collector, metricsServer, registry), nil
}
