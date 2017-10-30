package teams

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Delete removes a team.
func Delete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "teams", "delete")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Team(r.Context())

		if err := store.DeleteTeam(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to delete team",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to delete team"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
