package clean

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

type CompositeCleaner struct {
	cleaners map[string]Cleaner
}

func NewCompositeCLeaner(cleaners map[string]Cleaner) *CompositeCleaner {
	return &CompositeCleaner{
		cleaners: cleaners,
	}
}

func (c *CompositeCleaner) Clean(ctx context.Context) error {
	errs := []error{}

	for name, cleaner := range c.cleaners {
		slog.InfoContext(ctx, fmt.Sprintf("run %s", name))

		err := cleaner.Clean(ctx)
		if err != nil {
			slog.ErrorContext(
				ctx,
				fmt.Sprintf("[%s] failed to clean", name),
				slog.Any("err", err),
			)
		}
	}

	return errors.Join(errs...)
}
