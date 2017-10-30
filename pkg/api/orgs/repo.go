package orgs

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/json-iterator/go"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// RepoIndex retrieves all repos for the org.
func RepoIndex(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "repos/index")

	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.GetRepos(
			&model.ReposFilter{
				Org: session.Org(r.Context()),
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch repos",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to fetch repos"))
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

// RepoShow retrieves a repo for the org.
func RepoShow(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "repos/show")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Repo(r.Context())

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

// RepoDelete removes a repo for the org.
func RepoDelete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "repos/delete")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Repo(r.Context())

		if err := store.DeleteRepo(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to delete repo",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to delete repo"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
