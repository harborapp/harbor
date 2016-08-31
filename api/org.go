package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/router/middleware/session"
	"github.com/umschlag/umschlag-api/store"
)

// OrgIndex retrieves all available orgs.
func OrgIndex(c *gin.Context) {
	records, err := store.GetOrgs(
		c,
	)

	if err != nil {
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

// OrgShow retrieves a specific org.
func OrgShow(c *gin.Context) {
	record := session.Org(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// OrgDelete removes a specific org.
func OrgDelete(c *gin.Context) {
	record := session.Org(c)

	err := store.DeleteOrg(
		c,
		record,
	)

	if err != nil {
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
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted org",
		},
	)
}

// OrgUpdate updates an existing org.
func OrgUpdate(c *gin.Context) {
	record := session.Org(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind org data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind org data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateOrg(
		c,
		record,
	)

	if err != nil {
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

// OrgCreate creates a new org.
func OrgCreate(c *gin.Context) {
	record := &model.Org{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind org data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind org data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateOrg(
		c,
		record,
	)

	if err != nil {
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

// OrgUserIndex retrieves all users related to a org.
func OrgUserIndex(c *gin.Context) {
	records, err := store.GetOrgUsers(
		c,
		&model.OrgUserParams{
			Org: c.Param("org"),
		},
	)

	if err != nil {
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

// OrgUserAppend appends a user to a org.
func OrgUserAppend(c *gin.Context) {
	form := &model.OrgUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetOrgHasUser(
		c,
		form,
	)

	if assigned == true {
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

	err := store.CreateOrgUser(
		c,
		form,
	)

	if err != nil {
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

// OrgUserPerm updates the org team permission.
func OrgUserPerm(c *gin.Context) {
	form := &model.OrgUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetOrgHasUser(
		c,
		form,
	)

	if assigned == false {
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

	err := store.UpdateOrgUser(
		c,
		form,
	)

	if err != nil {
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

// OrgUserDelete deleted a user from a org
func OrgUserDelete(c *gin.Context) {
	form := &model.OrgUserParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetOrgHasUser(
		c,
		form,
	)

	if assigned == false {
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

	err := store.DeleteOrgUser(
		c,
		form,
	)

	if err != nil {
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

// OrgTeamIndex retrieves all teams related to a org.
func OrgTeamIndex(c *gin.Context) {
	records, err := store.GetOrgTeams(
		c,
		&model.OrgTeamParams{
			Org: c.Param("org"),
		},
	)

	if err != nil {
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

// OrgTeamAppend appends a team to a org.
func OrgTeamAppend(c *gin.Context) {
	form := &model.OrgTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetOrgHasTeam(
		c,
		form,
	)

	if assigned == true {
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

	err := store.CreateOrgTeam(
		c,
		form,
	)

	if err != nil {
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

// OrgTeamPerm updates the org team permission.
func OrgTeamPerm(c *gin.Context) {
	form := &model.OrgTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetOrgHasTeam(
		c,
		form,
	)

	if assigned == false {
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

	err := store.UpdateOrgTeam(
		c,
		form,
	)

	if err != nil {
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

// OrgTeamDelete deleted a team from a org
func OrgTeamDelete(c *gin.Context) {
	form := &model.OrgTeamParams{}

	if err := c.BindJSON(&form); err != nil {
		logrus.Warn("Failed to bind post data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind form data",
			},
		)

		c.Abort()
		return
	}

	assigned := store.GetOrgHasTeam(
		c,
		form,
	)

	if assigned == false {
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

	err := store.DeleteOrgTeam(
		c,
		form,
	)

	if err != nil {
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
