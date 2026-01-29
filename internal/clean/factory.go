package clean

import (
	"github.com/artarts36/docker-cleanup/internal/metrics"
	"github.com/docker/docker/client"
)

type Opts struct {
	Containers bool
	Images     bool
}

func CreateCleaner(cli *client.Client, opts Opts, metricsCollector metrics.Collector) Cleaner {
	cleaners := map[string]Cleaner{}

	if opts.Containers {
		cleaners["container-cleaner"] = NewContainerCleaner(cli, metricsCollector)
	}

	if opts.Images {
		cleaners["image-cleaner"] = NewImageCleaner(cli, metricsCollector)
	}

	return NewCompositeCLeaner(cleaners)
}
