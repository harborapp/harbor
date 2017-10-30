package auth

import (
	"io"
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/config"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/storage"
	"github.com/umschlag/umschlag-api/pkg/token"
)

const (
	loginBodyLimit int64 = 3 * 1024 * 1024
)

// Login represents the login handler.
func Login(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "auth", "login")

	return func(w http.ResponseWriter, r *http.Request) {
		auth := &model.Auth{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, loginBodyLimit)).Decode(auth); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to bind login",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to bind login"))
			return
		}

		user, err := store.GetUser(
			auth.Username,
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch user",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unauthorized("wrong username or password"))
			return
		}

		if err := user.MatchPassword(auth.Password); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to match password",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unauthorized("wrong username or password"))
			return
		}

		token := token.New(token.SessToken, user.Username)
		result, err := token.SignExpiring(user.Hash, config.Session.Expire)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to generate token",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unauthorized("wrong username or password"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsoniter.NewEncoder(w).Encode(result)
	}
}
