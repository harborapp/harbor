package orgs

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

// UserIndex retrieves all users related to an org.
func UserIndex(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "users/index")

	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.GetOrgUsers(
			&model.OrgUserParams{
				Org: chi.URLParam(r, "org"),
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch org users",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to fetch users"))
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

// UserPerm updates the org user permission.
func UserPerm(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "users/perm")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.OrgUserParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.UserBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if !store.GetOrgHasUser(form) {
			level.Warn(logger).Log(
				"msg", "user is not assigned",
			)

			fail.Error(w, fail.BadRequest("user is not assigned"))
			return
		}

		if err := store.UpdateOrgUser(form); err != nil {
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

// UserAppend appends an user to an org.
func UserAppend(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "users/append")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.OrgUserParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.UserBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if store.GetOrgHasUser(form) {
			level.Warn(logger).Log(
				"msg", "user is already appended",
			)

			fail.Error(w, fail.BadRequest("user is already appended"))
			return
		}

		if err := store.CreateOrgUser(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to append user",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to append user"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

// UserDelete deletes an user from an org.
func UserDelete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "users/delete")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.OrgUserParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.UserBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if !store.GetOrgHasUser(form) {
			level.Warn(logger).Log(
				"msg", "user is not assigned",
			)

			fail.Error(w, fail.BadRequest("user is not assigned"))
			return
		}

		if err := store.DeleteOrgUser(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to unlink user",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to unlink user"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
