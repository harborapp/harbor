package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/store"
)

// TeamIndex retrieves all available teams.
func TeamIndex(c *gin.Context) {
	records, err := store.GetTeams(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch teams. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch teams",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// TeamShow retrieves a specific team.
func TeamShow(c *gin.Context) {
	record := session.Team(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// TeamDelete removes a specific team.
func TeamDelete(c *gin.Context) {
	record := session.Team(c)

	err := store.DeleteTeam(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete team",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted team",
		},
	)
}

// TeamUpdate updates an existing team.
func TeamUpdate(c *gin.Context) {
	record := session.Team(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateTeam(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update team. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// TeamCreate creates a new user.
func TeamCreate(c *gin.Context) {
	record := &model.Team{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateTeam(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create team. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// TeamUserIndex retrieves all users related to a team.
func TeamUserIndex(c *gin.Context) {
	records, err := store.GetTeamUsers(
		c,
		&model.TeamUserParams{
			Team: c.Param("team"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch team users. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch users",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// TeamUserAppend appends a user to a team.
func TeamUserAppend(c *gin.Context) {
	form := &model.TeamUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasUser(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateTeamUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append team user. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended user",
		},
	)
}

// TeamUserPerm updates the org team permission.
func TeamUserPerm(c *gin.Context) {
	form := &model.TeamUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasUser(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateTeamUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to update permissions. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to update permissions",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully updated permissions",
		},
	)
}

// TeamUserDelete deleted a user from a team
func TeamUserDelete(c *gin.Context) {
	form := &model.TeamUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team user data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasUser(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "User is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteTeamUser(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team user. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked user",
		},
	)
}

// TeamOrgIndex retrieves all orgs related to a team.
func TeamOrgIndex(c *gin.Context) {
	records, err := store.GetTeamOrgs(
		c,
		&model.TeamOrgParams{
			Team: c.Param("team"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch team orgs. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch orgs",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// TeamOrgAppend appends a org to a team.
func TeamOrgAppend(c *gin.Context) {
	form := &model.TeamOrgParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team org data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team org data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasOrg(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Org is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateTeamOrg(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append team org. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append org",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended org",
		},
	)
}

// TeamOrgPerm updates the org team permission.
func TeamOrgPerm(c *gin.Context) {
	form := &model.TeamOrgParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team org data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team org data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasOrg(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Org is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateTeamOrg(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to update permissions. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to update permissions",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully updated permissions",
		},
	)
}

// TeamOrgDelete deleted a org from a team
func TeamOrgDelete(c *gin.Context) {
	form := &model.TeamOrgParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind team org data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind team org data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetTeamHasOrg(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Org is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteTeamOrg(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete team org. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink org",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked org",
		},
	)
}
