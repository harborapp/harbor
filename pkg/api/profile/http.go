package profile

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// NewHandler initializes the muxer for profile routes.
func NewHandler(store storage.Store, logger log.Logger) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(session.MustCurrent())

	mux.Get("/token", Token(store, logger))
	mux.Get("/self", Show(store, logger))
	mux.Put("/self", Update(store, logger))

	return mux
}
