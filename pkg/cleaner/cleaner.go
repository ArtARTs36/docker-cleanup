package cleaner

import (
	"github.com/docker/docker/client"

	"github.com/artarts36/docker-cleanup/internal/clean"
	"github.com/artarts36/docker-cleanup/internal/metrics"
)

type (
	Cleaner clean.Cleaner

	Opts struct {
		Containers bool
		Images     bool

		MetricsCollector MetricsCollector
	}
)

func New(cli *client.Client, opts Opts) Cleaner {
	if opts.MetricsCollector == nil {
		opts.MetricsCollector = metrics.NoopCollector{}
	}

	return clean.CreateCleaner(cli, clean.Opts{
		Containers: opts.Containers,
		Images:     opts.Images,
	}, opts.MetricsCollector)
}
