package teams

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// NewHandler initializes the muxer for teams routes.
func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(session.MustCurrent())
	mux.Use(session.MustTeams("display", store, logger))

	mux.Get("/", Index(store, logger))
	mux.With(session.SetTeam(store, logger)).Get("/{team}", Show(store, logger))
	mux.With(session.SetTeam(store, logger), session.MustTeams("delete", store, logger)).Delete("/{team}", Delete(store, logger))
	mux.With(session.SetTeam(store, logger), session.MustTeams("change", store, logger)).Put("/{team}", Update(store, logger))
	mux.With(session.MustTeams("change", store, logger)).Post("/", Create(store, logger))

	mux.Route("/{team}/users", func(t chi.Router) {
		t.Use(session.SetTeam(store, logger))

		t.With(session.MustTeamUsers("display", store, logger)).Get("/", UserIndex(store, logger))
		t.With(session.MustTeamUsers("change", store, logger)).Post("/", UserAppend(store, logger))
		t.With(session.MustTeamUsers("change", store, logger)).Put("/", UserPerm(store, logger))
		t.With(session.MustTeamUsers("change", store, logger)).Delete("/", UserDelete(store, logger))
	})

	mux.Route("/{team}/orgs", func(o chi.Router) {
		o.Use(session.SetTeam(store, logger))

		o.With(session.MustTeamOrgs("display", store, logger)).Get("/", OrgIndex(store, logger))
		o.With(session.MustTeamOrgs("change", store, logger)).Post("/", OrgAppend(store, logger))
		o.With(session.MustTeamOrgs("change", store, logger)).Put("/", OrgPerm(store, logger))
		o.With(session.MustTeamOrgs("change", store, logger)).Delete("/", OrgDelete(store, logger))
	})

	return mux
}
