package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// TagContextKey defines the context key that stores the tag.
	TagContextKey = "tag"
)

// Tag gets the tag from the context.
func Tag(c *gin.Context) *model.Tag {
	v, ok := c.Get(TagContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Tag)

	if !ok {
		return nil
	}

	return r
}

// SetTag injects the tag into the context.
func SetTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetTag(
			c,
			c.Param("tag"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find tag",
				},
			)

			c.Abort()
		} else {
			c.Set(TagContextKey, record)
			c.Next()
		}
	}
}

// MustTags validates the tags access.
func MustTags(action string) gin.HandlerFunc {
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
			case action == "display": // && user.Permission.DisplayTags:
				c.Next()
			case action == "change": // && user.Permission.ChangeTags:
				c.Next()
			case action == "delete": // && user.Permission.DeleteTags:
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
