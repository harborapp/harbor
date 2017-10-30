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
	// UserBodyLimit defines the maximum allowed POST body size.
	UserBodyLimit int64 = 3 * 1024 * 1024

	// UserContextKey defines the context key for the user context store.
	UserContextKey = &contextKey{"user"}
)

// User gets the user from the context.
func User(c context.Context) *model.User {
	v, ok := c.Value(UserContextKey).(*model.User)

	if !ok {
		return nil
	}

	return v
}

// SetUser injects the user into the context.
func SetUser(store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	logger = log.WithPrefix(logger, "session", "user")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			record, err := store.GetUser(
				chi.URLParam(r, "user"),
			)

			if err != nil {
				level.Warn(logger).Log(
					"msg", "failed to find user",
					"err", err,
				)

				fail.Error(w, fail.Cause(err).NotFound("failed to find user"))
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, record)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustUsers validates the users access.
func MustUsers(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowUserDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "delete":
				if allowUserDelete(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "update":
				if allowUserUpdate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "create":
				if allowUserCreate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowUserDisplay checks if the given user is allowed to display the resource.
func allowUserDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowUserDelete checks if the given user is allowed to delete the resource.
func allowUserDelete(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowUserUpdate checks if the given user is allowed to change the resource.
func allowUserUpdate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowUserCreate checks if the given user is allowed to change the resource.
func allowUserCreate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// MustUserTeams validates the user teams access.
func MustUserTeams(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowUserTeamDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "change":
				if allowUserTeamChange(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowUserTeamDisplay checks if the given user is allowed to display the resource.
func allowUserTeamDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowUserTeamChange checks if the given user is allowed to change the resource.
func allowUserTeamChange(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// MustUserOrgs validates the user orgs access.
func MustUserOrgs(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowUserOrgDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "change":
				if allowUserOrgChange(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowUserOrgDisplay checks if the given user is allowed to display the resource.
func allowUserOrgDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowUserOrgChange checks if the given user is allowed to change the resource.
func allowUserOrgChange(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}
