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
	// TeamBodyLimit defines the maximum allowed POST body size.
	TeamBodyLimit int64 = 3 * 1024 * 1024

	// TeamContextKey defines the context key for the team context store.
	TeamContextKey = &contextKey{"team"}
)

// Team gets the team from the context.
func Team(c context.Context) *model.Team {
	v, ok := c.Value(TeamContextKey).(*model.Team)

	if !ok {
		return nil
	}

	return v
}

// SetTeam injects the team into the context.
func SetTeam(store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	logger = log.WithPrefix(logger, "session", "team")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			record, err := store.GetTeam(
				chi.URLParam(r, "team"),
			)

			if err != nil {
				level.Warn(logger).Log(
					"msg", "failed to find team",
					"err", err,
				)

				fail.Error(w, fail.Cause(err).NotFound("failed to find team"))
				return
			}

			ctx := context.WithValue(r.Context(), TeamContextKey, record)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// MustTeams validates the teams access.
func MustTeams(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowTeamDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "delete":
				if allowTeamDelete(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "update":
				if allowTeamUpdate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "create":
				if allowTeamCreate(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowTeamDisplay checks if the given user is allowed to display the resource.
func allowTeamDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTeamDelete checks if the given user is allowed to delete the resource.
func allowTeamDelete(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTeamUpdate checks if the given user is allowed to change the resource.
func allowTeamUpdate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTeamCreate checks if the given user is allowed to change the resource.
func allowTeamCreate(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// MustTeamUsers validates the team users access.
func MustTeamUsers(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowTeamUserDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "change":
				if allowTeamUserChange(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowTeamUserDisplay checks if the given user is allowed to display the resource.
func allowTeamUserDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTeamUserChange checks if the given user is allowed to change the resource.
func allowTeamUserChange(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// MustTeamOrgs validates the team orgs access.
func MustTeamOrgs(action string, store storage.Store, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := Current(r.Context())

			if current.Admin {
				next.ServeHTTP(w, r)
				return
			}

			switch {
			case action == "display":
				if allowTeamOrgDisplay(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			case action == "change":
				if allowTeamOrgChange(r.Context(), store, logger) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fail.Error(w, fail.Forbidden("not allowed to access this resource"))
		})
	}
}

// allowTeamOrgDisplay checks if the given user is allowed to display the resource.
func allowTeamOrgDisplay(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}

// allowTeamOrgChange checks if the given user is allowed to change the resource.
func allowTeamOrgChange(ctx context.Context, store storage.Store, logger log.Logger) bool {
	// TODO(tboerger): add real implementation
	return false
}
