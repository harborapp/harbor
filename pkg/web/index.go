package web

import (
	"net/http"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/umschlag/umschlag-api/pkg/config"
	"github.com/umschlag/umschlag-api/pkg/storage"
	"github.com/umschlag/umschlag-api/pkg/templates"
)

// Index represents the index page.
func Index(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "web", "index")

	return func(w http.ResponseWriter, r *http.Request) {
		if err := templates.Load(logger).ExecuteTemplate(w, "index.html", indexVars()); err != nil {
			level.Warn(logger).Log(
				"msg", "failed to process index template",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unexpected())
			return
		}
	}
}

func indexVars() map[string]string {
	return map[string]string{
		"Root": config.Server.Root,
	}
}
