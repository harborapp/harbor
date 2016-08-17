package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/router/middleware/session"
	"github.com/umschlag/umschlag-api/store"
)

// TagIndex retrieves all available tags.
func TagIndex(c *gin.Context) {
	records, err := store.GetTags(
		c,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch tags",
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

// TagShow retrieves a specific tag.
func TagShow(c *gin.Context) {
	record := session.Tag(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// TagDelete removes a specific tag.
func TagDelete(c *gin.Context) {
	record := session.Tag(c)

	err := store.DeleteTag(
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
			"message": "Successfully deleted tag",
		},
	)
}
