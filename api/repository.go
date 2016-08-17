package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/router/middleware/session"
	"github.com/umschlag/umschlag-api/store"
)

// RepositoryIndex retrieves all available repositories.
func RepositoryIndex(c *gin.Context) {
	records, err := store.GetRepositories(
		c,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch repositories",
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

// RepositoryShow retrieves a specific repository.
func RepositoryShow(c *gin.Context) {
	record := session.Repository(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// RepositoryDelete removes a specific repository.
func RepositoryDelete(c *gin.Context) {
	record := session.Repository(c)

	err := store.DeleteRepository(
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
			"message": "Successfully deleted repository",
		},
	)
}
