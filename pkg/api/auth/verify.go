package auth

import (
	"encoding/base32"
	"net/http"
	"time"

	"github.com/codehack/fail"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/storage"
	"github.com/umschlag/umschlag-api/pkg/token"
)

// Verify is a handler to verify an JWT token.
func Verify(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "auth", "verify")

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			record *model.User
		)

		_, err := token.Direct(
			chi.URLParam(r, "token"),
			func(t *token.Token) ([]byte, error) {
				var (
					err error
				)

				record, err = store.GetUser(
					t.Text,
				)

				signingKey, _ := base32.StdEncoding.DecodeString(record.Hash)
				return signingKey, err
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "invalid token provided",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unauthorized("invalid token provided"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsoniter.NewEncoder(w).Encode(struct {
			Username  string    `json:"username"`
			CreatedAt time.Time `json:"created_at"`
		}{
			Username:  record.Username,
			CreatedAt: record.CreatedAt,
		})
	}
}
