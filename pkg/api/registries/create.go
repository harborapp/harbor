package registries

import (
	"io"
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Create creates a registry.
func Create(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "registries", "create")

	return func(w http.ResponseWriter, r *http.Request) {
		record := &model.Registry{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.RegistryBodyLimit)).Decode(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if err := store.CreateRegistry(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to create registry",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to create registry"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := jsoniter.NewEncoder(w).Encode(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to generate response",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to generate response"))
			return
		}
	}
}
