package metrics

import "context"

type Collector interface {
	ContainersCleaned(count int)
	ImagesCleaned(count int)
	Flush(ctx context.Context) error
}
