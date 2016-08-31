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
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowRepoDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowRepoChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowRepoDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowRepoDisplay checks if the given user is allowed to display the resource.
func allowRepoDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowRepoChange checks if the given user is allowed to change the resource.
func allowRepoChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowRepoDelete checks if the given user is allowed to delete the resource.
func allowRepoDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
