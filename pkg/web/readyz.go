package web

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Readyz is a simple ready check used by Docker and Kubernetes.
func Readyz(store storage.Store, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, http.StatusText(http.StatusOK))
	}
}
