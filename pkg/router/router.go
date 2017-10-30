package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/api/auth"
	"github.com/umschlag/umschlag-api/pkg/api/general"
	"github.com/umschlag/umschlag-api/pkg/api/orgs"
	"github.com/umschlag/umschlag-api/pkg/api/profile"
	"github.com/umschlag/umschlag-api/pkg/api/registries"
	"github.com/umschlag/umschlag-api/pkg/api/teams"
	"github.com/umschlag/umschlag-api/pkg/api/users"
	"github.com/umschlag/umschlag-api/pkg/assets"
	"github.com/umschlag/umschlag-api/pkg/config"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/header"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/prometheus"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/requests"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
	"github.com/umschlag/umschlag-api/pkg/web"
)

// Load initializes the routing of the application.
func Load(store storage.Store, logger log.Logger) http.Handler {
	mux := chi.NewRouter()

	mux.Use(requests.Requests(logger))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)

	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Use(session.SetCurrent(store))

	mux.Route(config.Server.Root, func(root chi.Router) {
		if config.Server.Prometheus {
			root.Get("/metrics", prometheus.Handler())
		}

		if config.Server.Pprof {
			root.Mount("/debug", middleware.Profiler())
		}

		root.Get("/", web.Index(store, logger))
		root.Get("/favicon.ico", web.Favicon(store, logger))

		root.Get("/healthz", web.Healthz(store, logger))
		root.Get("/readyz", web.Readyz(store, logger))

		root.Route("/api", func(base chi.Router) {
			base.Get("/", general.Index(store, logger))

			base.Mount("/auth", auth.NewHandler(store, logger))
			base.Mount("/profile", profile.NewHandler(store, logger))
			base.Mount("/registries", registries.NewHandler(store, logger))
			base.Mount("/orgs", orgs.NewHandler(store, logger))
			base.Mount("/users", users.NewHandler(store, logger))
			base.Mount("/teams", teams.NewHandler(store, logger))
		})

		root.Handle("/assets/*", static(logger))
		root.Handle("/storage/*", files(logger))
	})

	return mux
}

func static(logger log.Logger) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf(
			"%sassets",
			config.Server.Root,
		),
		http.FileServer(
			assets.Load(logger),
		),
	)
}

func files(logger log.Logger) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf(
			"%sstorage",
			config.Server.Root,
		),
		http.FileServer(
			http.Dir(config.Server.Storage),
		),
	)
}
