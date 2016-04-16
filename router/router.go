package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harborapp/harbor-api/assets"
	"github.com/harborapp/harbor-api/config"
	"github.com/harborapp/harbor-api/controller"
	"github.com/harborapp/harbor-api/router/middleware/context"
	"github.com/harborapp/harbor-api/router/middleware/header"
	"github.com/harborapp/harbor-api/router/middleware/logger"
	"github.com/harborapp/harbor-api/router/middleware/recovery"
	"github.com/harborapp/harbor-api/router/middleware/session"
	"github.com/harborapp/harbor-api/template"
	"github.com/harborapp/harbor-api/web"
)

// Load initializes the routing of the application.
func Load(cfg *config.Config, middleware ...gin.HandlerFunc) http.Handler {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.SetHTMLTemplate(
		template.Load(),
	)

	e.Use(middleware...)
	e.Use(context.SetLocation())
	e.Use(logger.SetLogger())
	e.Use(recovery.SetRecovery())
	e.Use(header.SetCache())
	e.Use(header.SetOptions())
	e.Use(header.SetSecure())
	e.Use(header.SetVersion())
	e.Use(session.SetCurrent())

	r := e.Group(cfg.Server.Root)
	{
		r.StaticFS(
			"/storage",
			gin.Dir(
				cfg.Server.Storage,
				false,
			),
		)

		r.StaticFS(
			"/assets",
			assets.Load(),
		)

		r.GET("/favicon.ico", web.GetFavicon)
		r.GET("", web.GetIndex)

		api := r.Group("/api")
		{
			api.GET("", controller.GetAPI)

			//
			// Profile
			//
			profile := api.Group("/profile")
			{
				profile.Use(session.MustCurrent())

				profile.GET("", controller.GetProfile)
				profile.PATCH("", controller.PatchProfile)
			}
		}
	}

	return e
}
