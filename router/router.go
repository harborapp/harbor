package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/api"
	"github.com/umschlag/umschlag-api/assets"
	"github.com/umschlag/umschlag-api/config"
	"github.com/umschlag/umschlag-api/router/middleware/header"
	"github.com/umschlag/umschlag-api/router/middleware/location"
	"github.com/umschlag/umschlag-api/router/middleware/logger"
	"github.com/umschlag/umschlag-api/router/middleware/recovery"
	"github.com/umschlag/umschlag-api/router/middleware/session"
	"github.com/umschlag/umschlag-api/router/middleware/store"
	"github.com/umschlag/umschlag-api/template"
	"github.com/umschlag/umschlag-api/web"
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

			//
			// Registries
			//
			registries := base.Group("/registries")
			{
				registries.Use(session.MustRegistries("display"))

				registries.GET("", api.RegistryIndex)
				registries.GET("/:registry", session.SetRegistry(), api.RegistryShow)
				registries.DELETE("/:registry", session.SetRegistry(), session.MustRegistries("delete"), api.RegistryDelete)
				registries.PATCH("/:registry", session.SetRegistry(), session.MustRegistries("change"), api.RegistryUpdate)
				registries.POST("", session.MustRegistries("change"), api.RegistryCreate)
			}

			//
			// Users
			//
			users := base.Group("/users")
			{
				users.Use(session.MustUsers("display"))

				users.GET("", api.UserIndex)
				users.GET("/:user", session.SetUser(), api.UserShow)
				users.DELETE("/:user", session.SetUser(), session.MustUsers("delete"), api.UserDelete)
				users.PATCH("/:user", session.SetUser(), session.MustUsers("change"), api.UserUpdate)
				users.POST("", session.MustUsers("change"), api.UserCreate)
			}

			userTeams := base.Group("/users/:user/teams")
			{
				userTeams.Use(session.MustTeams("change"))
				userTeams.Use(session.SetUser())

				userTeams.GET("", api.UserTeamIndex)
				userTeams.PATCH("", api.UserTeamAppend)
				userTeams.DELETE("", api.UserTeamDelete)
			}

			//
			// Teams
			//
			teams := base.Group("/teams")
			{
				teams.Use(session.MustTeams("display"))

				teams.GET("", api.TeamIndex)
				teams.GET("/:team", session.SetTeam(), api.TeamShow)
				teams.DELETE("/:team", session.SetTeam(), session.MustTeams("delete"), api.TeamDelete)
				teams.PATCH("/:team", session.SetTeam(), session.MustTeams("change"), api.TeamUpdate)
				teams.POST("", session.MustTeams("change"), api.TeamCreate)
			}

			teamUsers := base.Group("/teams/:team/users")
			{
				teamUsers.Use(session.MustTeams("change"))
				teamUsers.Use(session.SetTeam())

				teamUsers.GET("", api.TeamUserIndex)
				teamUsers.PATCH("", api.TeamUserAppend)
				teamUsers.DELETE("", api.TeamUserDelete)
			}

		}
	}

	return e
}
