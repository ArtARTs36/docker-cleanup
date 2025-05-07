package clean

import "github.com/docker/docker/client"

type Opts struct {
	Containers bool
	Images     bool
}

func CreateCleaner(cli *client.Client, opts Opts) Cleaner {
	cleaners := map[string]Cleaner{}

	if opts.Containers {
		cleaners["container-cleaner"] = NewContainerCleaner(cli)
	}

	if opts.Images {
		cleaners["image-cleaner"] = NewImageCleaner(cli)
	}

	return NewCompositeCLeaner(cleaners)
}
