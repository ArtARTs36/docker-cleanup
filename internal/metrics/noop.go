package metrics

import "context"

type NoopCollector struct{}

func (n NoopCollector) ContainersCleaned(_ int)       {}
func (n NoopCollector) ImagesCleaned(_ int)           {}
func (n NoopCollector) Flush(_ context.Context) error { return nil }
