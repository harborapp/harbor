package orgs

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// NewHandler initializes the muxer for orgs routes.
func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(session.MustCurrent())
	mux.Use(session.MustOrgs("display", store, logger))

	mux.Get("/", Index(store, logger))
	mux.With(session.SetOrg(store, logger)).Get("/{org}", Show(store, logger))
	mux.With(session.SetOrg(store, logger), session.MustOrgs("delete", store, logger)).Delete("/{org}", Delete(store, logger))
	mux.With(session.SetOrg(store, logger), session.MustOrgs("change", store, logger)).Put("/{org}", Update(store, logger))
	mux.With(session.MustOrgs("change", store, logger)).Post("/", Create(store, logger))

	mux.Route("/{org}/teams", func(t chi.Router) {
		t.Use(session.SetOrg(store, logger))

		t.With(session.MustOrgTeams("display", store, logger)).Get("/", TeamIndex(store, logger))
		t.With(session.MustOrgTeams("change", store, logger)).Post("/", TeamAppend(store, logger))
		t.With(session.MustOrgTeams("change", store, logger)).Put("/", TeamPerm(store, logger))
		t.With(session.MustOrgTeams("change", store, logger)).Delete("/", TeamDelete(store, logger))
	})

	mux.Route("/{org}/users", func(u chi.Router) {
		u.Use(session.SetOrg(store, logger))

		u.With(session.MustOrgUsers("display", store, logger)).Get("/", UserIndex(store, logger))
		u.With(session.MustOrgUsers("change", store, logger)).Post("/", UserAppend(store, logger))
		u.With(session.MustOrgUsers("change", store, logger)).Put("/", UserPerm(store, logger))
		u.With(session.MustOrgUsers("change", store, logger)).Delete("/", UserDelete(store, logger))
	})

	mux.Route("/{org}/repos", func(r chi.Router) {
		r.Use(session.SetOrg(store, logger))
		r.Use(session.MustRepos("display", store, logger))

		r.Get("/", RepoIndex(store, logger))
		r.With(session.SetRepo(store, logger)).Get("/{repo}", RepoShow(store, logger))
		r.With(session.SetRepo(store, logger), session.MustRepos("delete", store, logger)).Delete("/{repo}", RepoDelete(store, logger))
	})

	mux.Route("/{org}/repos/{repo}/tags", func(t chi.Router) {
		t.Use(session.SetOrg(store, logger))
		t.Use(session.SetRepo(store, logger))
		t.Use(session.MustTags("display", store, logger))

		t.Get("/", TagIndex(store, logger))
		t.With(session.SetTag(store, logger)).Get("/{tag}", TagShow(store, logger))
		t.With(session.SetTag(store, logger), session.MustTags("delete", store, logger)).Delete("/{tag}", TagDelete(store, logger))
	})

	return mux
}
