package router

import (
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/umschlag/umschlag-api/pkg/config"
	"github.com/umschlag/umschlag-api/pkg/middleware/header"
	"github.com/umschlag/umschlag-api/pkg/middleware/prometheus"
	"github.com/umschlag/umschlag-api/pkg/store"
	"github.com/umschlag/umschlag-api/pkg/swagger"
	"github.com/umschlag/umschlag-api/pkg/upload"
	"github.com/webhippie/fail"
)

// Server initializes the routing of the server.
func Server(cfg *config.Config, storage store.Store, uploads upload.Upload) http.Handler {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-Id"))

	mux.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Debug().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)

	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		root.Route("/api", func(base chi.Router) {
			base.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
				content, err := swagger.ReadFile("swagger.json")

				if err != nil {
					log.Error().
						Err(err).
						Msg("failed to read swagger.json")

					fail.ErrorJSON(w, fail.Unexpected())
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				io.WriteString(w, string(content))
			})

			if cfg.Server.Pprof {
				base.Mount("/debug", middleware.Profiler())
			}

			base.Handle("/storage/*", uploads.Handler(
				path.Join(
					cfg.Server.Root,
					"api",
					"storage",
				),
			))
		})
	})

	return mux
}

// Metrics initializes the routing of the metrics.
func Metrics(cfg *config.Config, storage store.Store, uploads upload.Upload) http.Handler {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-Id"))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)

	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route("/", func(root chi.Router) {
		root.Get("/metrics", prometheus.Handler(cfg.Metrics.Token))

		root.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			io.WriteString(w, http.StatusText(http.StatusOK))
		})

		root.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			io.WriteString(w, http.StatusText(http.StatusOK))
		})
	})

	return mux
}

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/go-kit/kit/log"
// 	"github.com/umschlag/umschlag-api/pkg/api/auth"
// 	"github.com/umschlag/umschlag-api/pkg/api/general"
// 	"github.com/umschlag/umschlag-api/pkg/api/orgs"
// 	"github.com/umschlag/umschlag-api/pkg/api/profile"
// 	"github.com/umschlag/umschlag-api/pkg/api/registries"
// 	"github.com/umschlag/umschlag-api/pkg/api/teams"
// 	"github.com/umschlag/umschlag-api/pkg/api/users"
// 	"github.com/umschlag/umschlag-api/pkg/assets"
// 	"github.com/umschlag/umschlag-api/pkg/config"
// 	"github.com/umschlag/umschlag-api/pkg/router/middleware/header"
// 	"github.com/umschlag/umschlag-api/pkg/router/middleware/prometheus"
// 	"github.com/umschlag/umschlag-api/pkg/router/middleware/requests"
// 	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
// 	"github.com/umschlag/umschlag-api/pkg/storage"
// 	"github.com/umschlag/umschlag-api/pkg/web"
// )

// // Load initializes the routing of the application.
// func Load(store storage.Store, logger log.Logger) http.Handler {
// 	mux := chi.NewRouter()

// 	mux.Use(requests.Requests(logger))

// 	mux.Use(middleware.Timeout(60 * time.Second))
// 	mux.Use(middleware.RealIP)

// 	mux.Use(header.Version)
// 	mux.Use(header.Cache)
// 	mux.Use(header.Secure)
// 	mux.Use(header.Options)

// 	mux.Use(session.SetCurrent(store))

// 	mux.Route(config.Server.Root, func(root chi.Router) {
// 		if config.Server.Prometheus {
// 			root.Get("/metrics", prometheus.Handler())
// 		}

// 		if config.Server.Pprof {
// 			root.Mount("/debug", middleware.Profiler())
// 		}

// 		root.Get("/", web.Index(store, logger))
// 		root.Get("/favicon.ico", web.Favicon(store, logger))

// 		root.Get("/healthz", web.Healthz(store, logger))
// 		root.Get("/readyz", web.Readyz(store, logger))

// 		root.Route("/api", func(base chi.Router) {
// 			base.Get("/", general.Index(store, logger))

// 			base.Mount("/auth", auth.NewHandler(store, logger))
// 			base.Mount("/profile", profile.NewHandler(store, logger))
// 			base.Mount("/registries", registries.NewHandler(store, logger))
// 			base.Mount("/orgs", orgs.NewHandler(store, logger))
// 			base.Mount("/users", users.NewHandler(store, logger))
// 			base.Mount("/teams", teams.NewHandler(store, logger))
// 		})

// 		root.Handle("/assets/*", static(logger))
// 		root.Handle("/storage/*", files(logger))
// 	})

// 	return mux
// }

// func static(logger log.Logger) http.Handler {
// 	return http.StripPrefix(
// 		fmt.Sprintf(
// 			"%sassets",
// 			config.Server.Root,
// 		),
// 		http.FileServer(
// 			assets.Load(logger),
// 		),
// 	)
// }

// func files(logger log.Logger) http.Handler {
// 	return http.StripPrefix(
// 		fmt.Sprintf(
// 			"%sstorage",
// 			config.Server.Root,
// 		),
// 		http.FileServer(
// 			http.Dir(config.Server.Storage),
// 		),
// 	)
// }
