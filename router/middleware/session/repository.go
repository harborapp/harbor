package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// RepositoryContextKey defines the context key that stores the repository.
	RepositoryContextKey = "repository"
)

// Repository gets the repository from the context.
func Repository(c *gin.Context) *model.Repository {
	v, ok := c.Get(RepositoryContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Repository)

	if !ok {
		return nil
	}

	return r
}

// SetRepository injects the repository into the context.
func SetRepository() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetRepository(
			c,
			c.Param("repo"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find repository",
				},
			)

			c.Abort()
		} else {
			c.Set(RepositoryContextKey, record)
			c.Next()
		}
	}
}

// MustRepositories validates the repositories access.
func MustRepositories(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := Current(c)

		if user == nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "You have to be authenticated",
				},
			)

			c.Abort()
		} else {
			switch {
			case action == "display": // && user.Permission.DisplayRepositories:
				c.Next()
			case action == "change": // && user.Permission.ChangeRepositories:
				c.Next()
			case action == "delete": // && user.Permission.DeleteRepositories:
				c.Next()
			default:
				c.JSON(
					http.StatusForbidden,
					gin.H{
						"status":  http.StatusForbidden,
						"message": "You are not authorized to request this resource",
					},
				)

				c.Abort()
			}
		}
	}
}
