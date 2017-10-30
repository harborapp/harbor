package auth

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Logout represents the logout handler.
func Logout(store storage.Store, logger log.Logger) http.HandlerFunc {
	// logger = log.WithPrefix(logger, "auth", "logout")

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
