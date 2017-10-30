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

// TagIndex retrieves all tags for the repo.
func TagIndex(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "tags/index")

	return func(w http.ResponseWriter, r *http.Request) {
		records, err := store.GetTags(
			&model.TagsFilter{
				Org:  session.Org(r.Context()),
				Repo: session.Repo(r.Context()),
			},
		)

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to fetch tags",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to fetch tags"))
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

// TagShow retrieves a tag for the repo.
func TagShow(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "tags/show")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Tag(r.Context())

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

// TagDelete removes a tag for the repo.
func TagDelete(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "orgs", "tags/delete")

	return func(w http.ResponseWriter, r *http.Request) {
		record := session.Tag(r.Context())

		if err := store.DeleteTag(record); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to delete tag",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).BadRequest("failed to delete tag"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
