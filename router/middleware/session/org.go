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
		org := Current(c)

		if org == nil {
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
			case action == "display": // && org.Permission.DisplayOrgs:
				c.Next()
			case action == "change": // && org.Permission.ChangeOrgs:
				c.Next()
			case action == "delete": // && org.Permission.DeleteOrgs:
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
