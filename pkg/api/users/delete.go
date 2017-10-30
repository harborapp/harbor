package users

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Delete removes an user.
func Delete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "delete")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.User(r.Context())

		if err := store.DeleteUser(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to delete user",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to delete user"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
