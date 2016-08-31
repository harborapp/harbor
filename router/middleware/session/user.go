package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// UserContextKey defines the context key that stores the user.
	UserContextKey = "user"
)

// User gets the user from the context.
func User(c *gin.Context) *model.User {
	v, ok := c.Get(UserContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.User)

	if !ok {
		return nil
	}

	return r
}

// SetUser injects the user into the context.
func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetUser(
			c,
			c.Param("user"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find user",
				},
			)

			c.Abort()
		} else {
			c.Set(UserContextKey, record)
			c.Next()
		}
	}
}

// MustUsers validates the users access.
func MustUsers(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowUserDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowUserChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowUserDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowUserDisplay checks if the given user is allowed to display the resource.
func allowUserDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowUserChange checks if the given user is allowed to change the resource.
func allowUserChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowUserDelete checks if the given user is allowed to delete the resource.
func allowUserDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustUserTeams validates the minecraft teams access.
func MustUserTeams(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowUserTeamDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowUserTeamChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowUserTeamDisplay checks if the given user is allowed to display the resource.
func allowUserTeamDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowUserTeamChange checks if the given user is allowed to change the resource.
func allowUserTeamChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustUsers validates the minecraft orgs access.
func MustUserOrgs(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowUserOrgDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowUserOrgChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowUserOrgDisplay checks if the given user is allowed to display the resource.
func allowUserOrgDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowUserOrgChange checks if the given user is allowed to change the resource.
func allowUserOrgChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
