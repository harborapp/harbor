package orgs

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Delete removes an org.
func Delete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "delete")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Org(r.Context())

		if err := store.DeleteOrg(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to delete org",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to delete org"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
