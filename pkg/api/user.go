package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/store"
)

// UserIndex retrieves all available users.
func UserIndex(c *gin.Context) {
	records, err := store.GetUsers(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch users. %s", err)

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

// UserShow retrieves a specific user.
func UserShow(c *gin.Context) {
	record := session.User(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// UserDelete removes a specific user.
func UserDelete(c *gin.Context) {
	record := session.User(c)

	err := store.DeleteUser(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete user",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted user",
		},
	)
}

// UserUpdate updates an existing user.
func UserUpdate(c *gin.Context) {
	record := session.User(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUser(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update user. %s", err)

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

// UserCreate creates a new user.
func UserCreate(c *gin.Context) {
	record := &model.User{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind user data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUser(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create user. %s", err)

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

// UserTeamIndex retrieves all teams related to a user.
func UserTeamIndex(c *gin.Context) {
	records, err := store.GetUserTeams(
		c,
		&model.UserTeamParams{
			User: c.Param("user"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch user teams. %s", err)

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

// UserTeamAppend appends a team to a user.
func UserTeamAppend(c *gin.Context) {
	form := &model.UserTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasTeam(
		c,
		form,
	)

	if assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Team is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUserTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append user team. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append team",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended team",
		},
	)
}

// UserTeamPerm updates the org team permission.
func UserTeamPerm(c *gin.Context) {
	form := &model.UserTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasTeam(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Team is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUserTeam(
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

// UserTeamDelete deleted a team from a user
func UserTeamDelete(c *gin.Context) {
	form := &model.UserTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user team data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user team data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasTeam(
		c,
		form,
	)

	if !assigned {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Team is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteUserTeam(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user team. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink team",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked team",
		},
	)
}

// UserOrgIndex retrieves all orgs related to a user.
func UserOrgIndex(c *gin.Context) {
	records, err := store.GetUserOrgs(
		c,
		&model.UserOrgParams{
			User: c.Param("user"),
		},
	)

	if err != nil {
		logrus.Warnf("Failed to fetch user orgs. %s", err)

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

// UserOrgAppend appends a org to a user.
func UserOrgAppend(c *gin.Context) {
	form := &model.UserOrgParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user org data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user org data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasOrg(
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

	err := store.CreateUserOrg(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to append user org. %s", err)

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

// UserOrgPerm updates the org team permission.
func UserOrgPerm(c *gin.Context) {
	form := &model.UserOrgParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user org data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user org data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasOrg(
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

	err := store.UpdateUserOrg(
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

// UserOrgDelete deleted a org from a user
func UserOrgDelete(c *gin.Context) {
	form := &model.UserOrgParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warnf("Failed to bind user org data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user org data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetUserHasOrg(
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

	err := store.DeleteUserOrg(
		c,
		form,
	)

	if err != nil {
		logrus.Warnf("Failed to delete user org. %s", err)

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
