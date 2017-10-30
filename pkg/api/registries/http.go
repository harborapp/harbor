package registries

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// NewHandler initializes the muxer for registries routes.
func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(session.MustCurrent())
	mux.Use(session.MustRegistries("display", store, logger))

	mux.Get("/", Index(store, logger))
	mux.With(session.SetRegistry(store, logger)).Get("/{registry}", Show(store, logger))
	mux.With(session.SetRegistry(store, logger), session.MustRegistries("delete", store, logger)).Delete("/{registry}", Delete(store, logger))
	mux.With(session.SetRegistry(store, logger), session.MustRegistries("change", store, logger)).Put("/{registry}", Update(store, logger))
	mux.With(session.MustRegistries("change", store, logger)).Post("/", Create(store, logger))

	return mux
}
