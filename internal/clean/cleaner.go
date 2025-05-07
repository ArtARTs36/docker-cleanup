package clean

import "context"

type Cleaner interface {
	Clean(ctx context.Context) error
}
