package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/api"
	"github.com/umschlag/umschlag-api/assets"
	"github.com/umschlag/umschlag-api/config"
	"github.com/umschlag/umschlag-api/router/middleware/header"
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
				registries.Use(session.MustCurrent())
				registries.Use(session.MustRegistries("display"))

				registries.GET("", api.RegistryIndex)
				registries.GET("/:registry", session.SetRegistry(), api.RegistryShow)
				registries.DELETE("/:registry", session.SetRegistry(), session.MustRegistries("delete"), api.RegistryDelete)
				registries.PATCH("/:registry", session.SetRegistry(), session.MustRegistries("change"), api.RegistryUpdate)
				registries.POST("", session.MustRegistries("change"), api.RegistryCreate)
			}

			//
			// Tags
			//
			tags := base.Group("/orgs/:org/repos/:repo/tags")
			{
				tags.Use(session.SetOrg())
				tags.Use(session.SetRepo())
				tags.Use(session.MustTags("display"))

				tags.GET("", api.TagIndex)
				tags.GET("/:tag", session.SetTag(), api.TagShow)
				tags.DELETE("/:tag", session.SetTag(), session.MustTags("delete"), api.TagDelete)
			}

			//
			// Repos
			//
			repos := base.Group("/orgs/:org/repos")
			{
				repos.Use(session.SetOrg())
				repos.Use(session.MustRepos("display"))

				repos.GET("", api.RepoIndex)
				repos.GET("/:repo", session.SetRepo(), api.RepoShow)
				repos.DELETE("/:repo", session.SetRepo(), session.MustRepos("delete"), api.RepoDelete)
			}

			//
			// Orgs
			//
			orgs := base.Group("/orgs")
			{
				orgs.Use(session.MustOrgs("display"))

				orgs.GET("", api.OrgIndex)
				orgs.GET("/:org", session.SetOrg(), api.OrgShow)
				orgs.DELETE("/:org", session.SetOrg(), session.MustOrgs("delete"), api.OrgDelete)
				orgs.PATCH("/:org", session.SetOrg(), session.MustOrgs("change"), api.OrgUpdate)
				orgs.POST("", session.MustOrgs("change"), api.OrgCreate)
			}

			orgTeams := base.Group("/orgs/:org/teams")
			{
				orgTeams.Use(session.MustCurrent())
				orgTeams.Use(session.SetOrg())

				orgTeams.GET("", session.MustOrgTeams("display"), api.OrgTeamIndex)
				orgTeams.POST("", session.MustOrgTeams("change"), api.OrgTeamAppend)
				orgTeams.PATCH("", session.MustOrgTeams("change"), api.OrgTeamPerm)
				orgTeams.DELETE("", session.MustOrgTeams("change"), api.OrgTeamDelete)
			}

			orgUsers := base.Group("/orgs/:org/users")
			{
				orgUsers.Use(session.MustCurrent())
				orgUsers.Use(session.SetOrg())

				orgUsers.GET("", session.MustOrgUsers("display"), api.OrgUserIndex)
				orgUsers.POST("", session.MustOrgUsers("change"), api.OrgUserAppend)
				orgUsers.PATCH("", session.MustOrgUsers("change"), api.OrgUserPerm)
				orgUsers.DELETE("", session.MustOrgUsers("change"), api.OrgUserDelete)
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
				userTeams.Use(session.MustCurrent())
				userTeams.Use(session.SetUser())

				userTeams.GET("", session.MustUserTeams("display"), api.UserTeamIndex)
				userTeams.POST("", session.MustUserTeams("change"), api.UserTeamAppend)
				userTeams.PATCH("", session.MustUserTeams("change"), api.UserTeamPerm)
				userTeams.DELETE("", session.MustUserTeams("change"), api.UserTeamDelete)
			}

			userOrgs := base.Group("/users/:user/orgs")
			{
				userOrgs.Use(session.MustCurrent())
				userOrgs.Use(session.SetUser())

				userOrgs.GET("", session.MustUserOrgs("display"), api.UserOrgIndex)
				userOrgs.POST("", session.MustUserOrgs("change"), api.UserOrgAppend)
				userOrgs.PATCH("", session.MustUserOrgs("change"), api.UserOrgPerm)
				userOrgs.DELETE("", session.MustUserOrgs("change"), api.UserOrgDelete)
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
				teamUsers.Use(session.MustCurrent())
				teamUsers.Use(session.SetTeam())

				teamUsers.GET("", session.MustTeamUsers("display"), api.TeamUserIndex)
				teamUsers.POST("", session.MustTeamUsers("change"), api.TeamUserAppend)
				teamUsers.PATCH("", session.MustTeamUsers("change"), api.TeamUserPerm)
				teamUsers.DELETE("", session.MustTeamUsers("change"), api.TeamUserDelete)
			}

			teamOrgs := base.Group("/teams/:team/orgs")
			{
				teamOrgs.Use(session.MustCurrent())
				teamOrgs.Use(session.SetTeam())

				teamOrgs.GET("", session.MustTeamOrgs("display"), api.TeamOrgIndex)
				teamOrgs.POST("", session.MustTeamOrgs("change"), api.TeamOrgAppend)
				teamOrgs.PATCH("", session.MustTeamOrgs("change"), api.TeamOrgPerm)
				teamOrgs.DELETE("", session.MustTeamOrgs("change"), api.TeamOrgDelete)
			}
		}
	}

	return e
}
