package distribution

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackspirou/syscerts"
)

// Repos represents a collection of docker distribution repos.
type Repos []*Repo

// Repo represents a repository from the docker distribution.
type Repo struct {
	Base string `json:"-"`

	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// URL returns the URL for the image details API endpoint.
func (u *Repo) URL() string {
	return strings.Join(
		[]string{
			u.Base,
			"v2",
			u.Name,
			"tags",
			"list",
		},
		"/",
	)
}

// Fetch retrieves more records from the distribution image API.
func (u *Repo) Fetch() error {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				RootCAs:            syscerts.SystemRootsPool(),
			},
		},
	}

	res, err := client.Get(u.URL())

	if err != nil {
		return fmt.Errorf("Failed to fetch repo content. %s", err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(u); err != nil {
		return fmt.Errorf("Failed to parse repo content. %s", err)
	}

	return nil
}
