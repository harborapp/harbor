package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// NamespaceContextKey defines the context key that stores the namespace.
	NamespaceContextKey = "namespace"
)

// Namespace gets the namespace from the context.
func Namespace(c *gin.Context) *model.Namespace {
	v, ok := c.Get(NamespaceContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Namespace)

	if !ok {
		return nil
	}

	return r
}

// SetNamespace injects the namespace into the context.
func SetNamespace() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetNamespace(
			c,
			c.Param("namespace"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find namespace",
				},
			)

			c.Abort()
		} else {
			c.Set(NamespaceContextKey, record)
			c.Next()
		}
	}
}

// MustNamespaces validates the namespaces access.
func MustNamespaces(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		namespace := Current(c)

		if namespace == nil {
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
			case action == "display": // && namespace.Permission.DisplayNamespaces:
				c.Next()
			case action == "change": // && namespace.Permission.ChangeNamespaces:
				c.Next()
			case action == "delete": // && namespace.Permission.DeleteNamespaces:
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
