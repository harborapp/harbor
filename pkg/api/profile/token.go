package profile

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
	"github.com/umschlag/umschlag-api/pkg/token"
)

// Token displays the user token.
func Token(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "profile", "token")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Current(r.Context())

		token := token.New(token.UserToken, record.Username)
		result, err := token.SignUnlimited(record.Hash)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to generate token",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to generate token"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := jsoniter.NewEncoder(w).Encode(result); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to generate response",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to generate response"))
			return
		}
	}
}
