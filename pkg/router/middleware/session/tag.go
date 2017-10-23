package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/store"
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
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowTagDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowTagChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowTagDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowTagDisplay checks if the given user is allowed to display the resource.
func allowTagDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTagChange checks if the given user is allowed to change the resource.
func allowTagChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowTagDelete checks if the given user is allowed to delete the resource.
func allowTagDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
