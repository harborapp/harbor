package session

import (
	"context"
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

var (
	// RepoContextKey defines the context key for the repo context store.
	RepoContextKey = &contextKey{"repo"}
)

// Repo gets the repo from the context.
func Repo(c context.Context) *model.Repo {
	v, ok := c.Value(RepoContextKey).(*model.Repo)

	if !ok {
		return nil
	}

	return v
}

// SetRepo injects the repo into the context.
func SetRepo(store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	logger = log.WithPrefix(logger, "session", "repo")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			record, err := store.GetRepo(
				chi.URLParam(r, "repo"),
			)

			if err != nil {
				level.Warn(logger).Log(
					"msg", "failed to find repo",
					"err", err,
				)

				fail.Error(w, fail.Cause(err).NotFound("failed to find repo"))
				return
			}

			ctx := context.WithValue(r.Context(), RepoContextKey, record)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustRepos validates the repos access.
func MustRepos(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowRepoDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "delete":
				if allowRepoDelete(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "update":
				if allowRepoUpdate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "create":
				if allowRepoCreate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowRepoDisplay checks if the given user is allowed to display the resource.
func allowRepoDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowRepoDelete checks if the given user is allowed to delete the resource.
func allowRepoDelete(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowRepoUpdate checks if the given user is allowed to change the resource.
func allowRepoUpdate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowRepoCreate checks if the given user is allowed to change the resource.
func allowRepoCreate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}
