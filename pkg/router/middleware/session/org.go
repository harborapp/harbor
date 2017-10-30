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
	// OrgBodyLimit defines the maximum allowed POST body size.
	OrgBodyLimit int64 = 3 * 1024 * 1024

	// OrgContextKey defines the context key for the org context store.
	OrgContextKey = &contextKey{"org"}
)

// Org gets the org from the context.
func Org(c context.Context) *model.Org {
	v, ok := c.Value(OrgContextKey).(*model.Org)

	if !ok {
		return nil
	}

	return v
}

// SetOrg injects the org into the context.
func SetOrg(store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	logger = log.WithPrefix(logger, "session", "org")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			record, err := store.GetOrg(
				chi.URLParam(r, "org"),
			)

			if err != nil {
				level.Warn(logger).Log(
					"msg", "failed to find org",
					"err", err,
				)

				fail.Error(w, fail.Cause(err).NotFound("failed to find org"))
				return
			}

			ctx := context.WithValue(r.Context(), OrgContextKey, record)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustOrgs validates the orgs access.
func MustOrgs(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowOrgDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "delete":
				if allowOrgDelete(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "update":
				if allowOrgUpdate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "create":
				if allowOrgCreate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowOrgDisplay checks if the given user is allowed to display the resource.
func allowOrgDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowOrgDelete checks if the given user is allowed to delete the resource.
func allowOrgDelete(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowOrgUpdate checks if the given user is allowed to change the resource.
func allowOrgUpdate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowOrgCreate checks if the given user is allowed to change the resource.
func allowOrgCreate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// MustOrgUsers validates the org users access.
func MustOrgUsers(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowOrgUserDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "change":
				if allowOrgUserChange(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowOrgUserDisplay checks if the given user is allowed to display the resource.
func allowOrgUserDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowOrgUserChange checks if the given user is allowed to change the resource.
func allowOrgUserChange(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// MustOrgTeams validates the org teams access.
func MustOrgTeams(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowOrgTeamDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "change":
				if allowOrgTeamChange(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowOrgTeamDisplay checks if the given user is allowed to display the resource.
func allowOrgTeamDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowOrgTeamChange checks if the given user is allowed to change the resource.
func allowOrgTeamChange(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}
