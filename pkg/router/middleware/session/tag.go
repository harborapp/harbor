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
	// TagContextKey defines the context key for the tag context store.
	TagContextKey = &contextKey{"tag"}
)

// Tag gets the tag from the context.
func Tag(c context.Context) *model.Tag {
	v, ok := c.Value(TagContextKey).(*model.Tag)

	if !ok {
		return nil
	}

	return v
}

// SetTag injects the tag into the context.
func SetTag(store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	logger = log.WithPrefix(logger, "session", "tag")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			record, err := store.GetTag(
				chi.URLParam(r, "tag"),
			)

			if err != nil {
				level.Warn(logger).Log(
					"msg", "failed to find tag",
					"err", err,
				)

				fail.Error(w, fail.Cause(err).NotFound("failed to find tag"))
				return
			}

			ctx := context.WithValue(r.Context(), TagContextKey, record)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustTags validates the tags access.
func MustTags(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowTagDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "delete":
				if allowTagDelete(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "update":
				if allowTagUpdate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "create":
				if allowTagCreate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowTagDisplay checks if the given user is allowed to display the resource.
func allowTagDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTagDelete checks if the given user is allowed to delete the resource.
func allowTagDelete(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTagUpdate checks if the given user is allowed to change the resource.
func allowTagUpdate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTagCreate checks if the given user is allowed to change the resource.
func allowTagCreate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}
