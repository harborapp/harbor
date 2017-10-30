package teams

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Index retrieves all teams.
func Index(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "teams", "index")

	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.GetTeams()

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch teams",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to fetch teams"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := jsoniter.NewEncoder(w).Encode(records); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to generate response",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to generate response"))
			return
		}
	}
}
