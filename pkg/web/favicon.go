package web

import (
	"net/http"
	"time"

	"github.com/codehack/fail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/umschlag/umschlag-api/pkg/assets"
	"github.com/umschlag/umschlag-api/pkg/storage"
)

// Favicon represents the favicon.
func Favicon(store storage.Store, logger log.Logger) http.HandlerFunc {
	logger = log.WithPrefix(logger, "web", "favicon")

	return func(w http.ResponseWriter, r *http.Request) {
		file, err := assets.Load(logger).Open("images/favicon.ico")

		if err != nil {
			level.Warn(logger).Log(
				"msg", "failed to retrieve favicon",
				"err", err,
			)

			fail.Error(w, fail.Cause(err).Unexpected())
			return
		}

		http.ServeContent(
			w,
			r,
			"favicon.ico",
			time.Now(),
			file,
		)
	}
}
