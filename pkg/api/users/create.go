package users

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

// Create creates an user.
func Create(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "create")

	return func(w http.ResponseWriter, r *http.Request) {
		record := &model.User{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.UserBodyLimit)).Decode(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if err := store.CreateUser(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to create user",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to create user"))
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
