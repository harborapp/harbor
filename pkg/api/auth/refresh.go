package auth

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/config"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
	"github.com/umschlag/umschlag-api/pkg/token"
)

// Refresh represents the refresh handler.
func Refresh(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "auth", "refresh")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Current(r.Context())

		token := token.New(token.SessToken, record.Username)
		result, err := token.SignExpiring(record.Hash, config.Session.Expire)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to refresh token",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unauthorized("failed to refresh token"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsoniter.NewEncoder(w).Encode(result)
	}
}
