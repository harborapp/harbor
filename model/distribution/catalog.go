package distribution

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackspirou/syscerts"
	"github.com/peterhellberg/link"
)

// Catalog represents the docker distribution catalog response.
type Catalog struct {
	Base string `json:"-"`
	Next string `json:"-"`

	Repos []string `json:"repositories"`
}

// URL returns the URL for the next catalog iteration.
func (u *Catalog) URL() string {
	return strings.Join(
		[]string{
			u.Base,
			u.Next,
		},
		"",
	)
}

// Over just checks if we need to iterate one more time the catalog.
func (u *Catalog) Over() bool {
	return u.Next == ""
}

// Fetch retrieves more records from the distribution catalog API.
func (u *Catalog) Fetch() error {
	var (
		temp = &Catalog{}
	)

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
		return fmt.Errorf("Failed to fetch catalog content. %s", err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(temp); err != nil {
		return fmt.Errorf("Failed to parse catalog content. %s", err)
	}

	u.Repos = append(
		u.Repos,
		temp.Repos...,
	)

	u.Next = ""

	for _, l := range link.ParseResponse(res) {
		if l.Rel == "next" {
			u.Next = l.URI
		}
	}

	return nil
}
