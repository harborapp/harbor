package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harborapp/harbor-api/api"
	"github.com/harborapp/harbor-api/assets"
	"github.com/harborapp/harbor-api/config"
	"github.com/harborapp/harbor-api/router/middleware/header"
	"github.com/harborapp/harbor-api/router/middleware/location"
	"github.com/harborapp/harbor-api/router/middleware/logger"
	"github.com/harborapp/harbor-api/router/middleware/recovery"
	"github.com/harborapp/harbor-api/router/middleware/session"
	"github.com/harborapp/harbor-api/router/middleware/store"
	"github.com/harborapp/harbor-api/template"
	"github.com/harborapp/harbor-api/web"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.SetHTMLTemplate(
		template.Load(),
	)

	e.Use(middleware...)
	e.Use(logger.SetLogger())
	e.Use(recovery.SetRecovery())
	e.Use(location.SetLocation())
	e.Use(store.SetStore())
	e.Use(header.SetCache())
	e.Use(header.SetOptions())
	e.Use(header.SetSecure())
	e.Use(header.SetVersion())
	e.Use(session.SetCurrent())

	root := e.Group(config.Server.Root)
	{
		root.StaticFS(
			"/storage",
			gin.Dir(
				config.Server.Storage,
				false,
			),
		)

		root.StaticFS(
			"/assets",
			assets.Load(),
		)

		root.GET("/favicon.ico", web.Favicon)
		root.GET("", web.Index)

		base := root.Group("/api")
		{
			base.GET("", api.IndexInfo)

			//
			// Auth
			//
			auth := base.Group("/auth")
			{
				auth.GET("/logout", session.MustCurrent(), api.AuthLogout)
				auth.GET("/refresh", session.MustCurrent(), api.AuthRefresh)
				auth.POST("/login", session.MustNobody(), api.AuthLogin)
			}

			//
			// Profile
			//
			profile := base.Group("/profile")
			{
				profile.Use(session.MustCurrent())

				profile.GET("/token", api.ProfileToken)
				profile.GET("/self", api.ProfileShow)
				profile.PATCH("/self", api.ProfileUpdate)
			}
		}
	}

	return e
}
