package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// RepoContextKey defines the context key that stores the repo.
	RepoContextKey = "repo"
)

// Repo gets the repo from the context.
func Repo(c *gin.Context) *model.Repo {
	v, ok := c.Get(RepoContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Repo)

	if !ok {
		return nil
	}

	return r
}

// SetRepo injects the repo into the context.
func SetRepo() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetRepo(
			c,
			c.Param("repo"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find repo",
				},
			)

			c.Abort()
		} else {
			c.Set(RepoContextKey, record)
			c.Next()
		}
	}
}

// MustRepos validates the repos access.
func MustRepos(action string) gin.HandlerFunc {
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
			case action == "display": // && user.Permission.DisplayRepos:
				c.Next()
			case action == "change": // && user.Permission.ChangeRepos:
				c.Next()
			case action == "delete": // && user.Permission.DeleteRepos:
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
