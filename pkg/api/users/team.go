package users

import (
	"io"
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// TeamIndex retrieves all teams related to an user.
func TeamIndex(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "teams/index")

	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.GetUserTeams(
			&model.UserTeamParams{
				User: chi.URLParam(r, "user"),
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch user teams",
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

// TeamPerm updates the user team permission.
func TeamPerm(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "teams/perm")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.UserTeamParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.TeamBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if !store.GetUserHasTeam(form) {
			level.Warn(logger).Log(
				"msg", "team is not assigned",
			)

			fail.Error(w, fail.BadRequest("team is not assigned"))
			return
		}

		if err := store.UpdateUserTeam(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to update permissions",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to update permissions"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

// TeamAppend appends a team to an user.
func TeamAppend(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "teams/append")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.UserTeamParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.TeamBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if store.GetUserHasTeam(form) {
			level.Warn(logger).Log(
				"msg", "team is already appended",
			)

			fail.Error(w, fail.BadRequest("team is already appended"))
			return
		}

		if err := store.CreateUserTeam(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to append team",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to append team"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

// TeamDelete deletes a team from an user.
func TeamDelete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "teams/delete")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.UserTeamParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.TeamBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if !store.GetUserHasTeam(form) {
			level.Warn(logger).Log(
				"msg", "team is not assigned",
			)

			fail.Error(w, fail.BadRequest("team is not assigned"))
			return
		}

		if err := store.DeleteUserTeam(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to unlink team",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to unlink team"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
