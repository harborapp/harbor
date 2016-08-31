package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// OrgContextKey defines the context key that stores the org.
	OrgContextKey = "org"
)

// Org gets the org from the context.
func Org(c *gin.Context) *model.Org {
	v, ok := c.Get(OrgContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Org)

	if !ok {
		return nil
	}

	return r
}

// SetOrg injects the org into the context.
func SetOrg() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetOrg(
			c,
			c.Param("org"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find org",
				},
			)

			c.Abort()
		} else {
			c.Set(OrgContextKey, record)
			c.Next()
		}
	}
}

// MustOrgs validates the orgs access.
func MustOrgs(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowOrgDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowOrgChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowOrgDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowOrgDisplay checks if the given user is allowed to display the resource.
func allowOrgDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowOrgChange checks if the given user is allowed to change the resource.
func allowOrgChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowOrgDelete checks if the given user is allowed to delete the resource.
func allowOrgDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustOrgUsers validates the minecraft builds access.
func MustOrgUsers(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowOrgUserDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowOrgUserChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowOrgUserDisplay checks if the given user is allowed to display the resource.
func allowOrgUserDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowOrgUserChange checks if the given user is allowed to change the resource.
func allowOrgUserChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// MustOrgTeams validates the minecraft builds access.
func MustOrgTeams(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowOrgTeamDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowOrgTeamChange(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowOrgTeamDisplay checks if the given user is allowed to display the resource.
func allowOrgTeamDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowOrgTeamChange checks if the given user is allowed to change the resource.
func allowOrgTeamChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
