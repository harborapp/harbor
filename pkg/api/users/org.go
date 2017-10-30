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

// OrgIndex retrieves all orgs related to an user.
func OrgIndex(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "orgs/index")

	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.GetUserOrgs(
			&model.UserOrgParams{
				User: chi.URLParam(r, "user"),
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch user orgs",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to fetch orgs"))
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

// OrgPerm updates the user org permission.
func OrgPerm(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "orgs/perm")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.UserOrgParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.OrgBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if !store.GetUserHasOrg(form) {
			level.Warn(logger).Log(
				"msg", "org is not assigned",
			)

			fail.Error(w, fail.BadRequest("org is not assigned"))
			return
		}

		if err := store.UpdateUserOrg(form); err != nil {
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

// OrgAppend appends an org to an user.
func OrgAppend(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "orgs/append")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.UserOrgParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.OrgBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if store.GetUserHasOrg(form) {
			level.Warn(logger).Log(
				"msg", "org is already appended",
			)

			fail.Error(w, fail.BadRequest("org is already appended"))
			return
		}

		if err := store.CreateUserOrg(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to append org",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to append org"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

// OrgDelete deletes an org from an user.
func OrgDelete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "users", "orgs/delete")

	return func(w http.ResponseWriter, r *http.Request) {
		form := &model.UserOrgParams{}

		if err := jsoniter.NewDecoder(io.LimitReader(r.Body, session.OrgBodyLimit)).Decode(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to parse request",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to parse request"))
			return
		}

		if !store.GetUserHasOrg(form) {
			level.Warn(logger).Log(
				"msg", "org is not assigned",
			)

			fail.Error(w, fail.BadRequest("org is not assigned"))
			return
		}

		if err := store.DeleteUserOrg(form); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to unlink org",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to unlink org"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
