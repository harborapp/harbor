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
	// RegistryBodyLimit defines the maximum allowed POST body size.
	RegistryBodyLimit int64 = 3 * 1024 * 1024

	// RegistryContextKey defines the context key for the registry context store.
	RegistryContextKey = &contextKey{"registry"}
)

// Registry gets the registry from the context.
func Registry(c context.Context) *model.Registry {
	v, ok := c.Value(RegistryContextKey).(*model.Registry)

	if !ok {
		return nil
	}

	return v
}

// SetRegistry injects the registry into the context.
func SetRegistry(store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	logger = log.WithPrefix(logger, "session", "registry")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			record, err := store.GetRegistry(
				chi.URLParam(r, "registry"),
			)

			if err != nil {
				level.Warn(logger).Log(
					"msg", "failed to find registry",
					"err", err,
				)

				fail.Error(w, fail.Cause(err).NotFound("failed to find registry"))
				return
			}

			ctx := context.WithValue(r.Context(), RegistryContextKey, record)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustRegistries validates the registries access.
func MustRegistries(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowRegistryDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "delete":
				if allowRegistryDelete(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "update":
				if allowRegistryUpdate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "create":
				if allowRegistryCreate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowRegistryDisplay checks if the given user is allowed to display the resource.
func allowRegistryDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowRegistryDelete checks if the given user is allowed to delete the resource.
func allowRegistryDelete(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowRegistryUpdate checks if the given user is allowed to change the resource.
func allowRegistryUpdate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowRegistryCreate checks if the given user is allowed to change the resource.
func allowRegistryCreate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}
