package distribution

import (
	"fmt"
)

const (
	CatalogPageSize = 100
)

// Load initializes and fetches the docker distribution catalog.
func Load(base string) (Repos, error) {
	var (
		repos = Repos{}

		catalog = &Catalog{
			Base: base,
			Next: fmt.Sprintf("/v2/_catalog?n=%d", CatalogPageSize),
		}
	)

	for {
		if catalog.Over() {
			break
		}

		if err := catalog.Fetch(); err != nil {
			return nil, err
		}
	}

	for _, name := range catalog.Repos {
		repo := &Repo{
			Base: base,
			Name: name,
		}

		if err := repo.Fetch(); err != nil {
			return nil, err
		}

		repos = append(
			repos,
			repo,
		)
	}

	return repos, nil
}
