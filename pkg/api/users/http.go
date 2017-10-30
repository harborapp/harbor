package users

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// NewHandler initializes the muxer for users routes.
func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(session.MustCurrent())
	mux.Use(session.MustUsers("display", store, logger))

	mux.Get("/", Index(store, logger))
	mux.With(session.SetUser(store, logger)).Get("/{user}", Show(store, logger))
	mux.With(session.SetUser(store, logger), session.MustUsers("delete", store, logger)).Delete("/{user}", Delete(store, logger))
	mux.With(session.SetUser(store, logger), session.MustUsers("change", store, logger)).Put("/{user}", Update(store, logger))
	mux.With(session.MustUsers("change", store, logger)).Post("/", Create(store, logger))

	mux.Route("/{user}/teams", func(t chi.Router) {
		t.Use(session.SetUser(store, logger))

		t.With(session.MustUserTeams("display", store, logger)).Get("/", TeamIndex(store, logger))
		t.With(session.MustUserTeams("change", store, logger)).Post("/", TeamAppend(store, logger))
		t.With(session.MustUserTeams("change", store, logger)).Put("/", TeamPerm(store, logger))
		t.With(session.MustUserTeams("change", store, logger)).Delete("/", TeamDelete(store, logger))
	})

	mux.Route("/{user}/orgs", func(o chi.Router) {
		o.Use(session.SetUser(store, logger))

		o.With(session.MustUserOrgs("display", store, logger)).Get("/", OrgIndex(store, logger))
		o.With(session.MustUserOrgs("change", store, logger)).Post("/", OrgAppend(store, logger))
		o.With(session.MustUserOrgs("change", store, logger)).Put("/", OrgPerm(store, logger))
		o.With(session.MustUserOrgs("change", store, logger)).Delete("/", OrgDelete(store, logger))
	})

	return mux
}
