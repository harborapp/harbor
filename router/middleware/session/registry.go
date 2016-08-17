package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/model"
	"github.com/umschlag/umschlag-api/store"
)

const (
	// RegistryContextKey defines the context key that stores the registry.
	RegistryContextKey = "registry"
)

// Registry gets the registry from the context.
func Registry(c *gin.Context) *model.Registry {
	v, ok := c.Get(RegistryContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Registry)

	if !ok {
		return nil
	}

	return r
}

// SetRegistry injects the registry into the context.
func SetRegistry() gin.HandlerFunc {
	return func(c *gin.Context) {
		record, res := store.GetRegistry(
			c,
			c.Param("registry"),
		)

		if res.Error != nil || res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find registry",
				},
			)

			c.Abort()
		} else {
			c.Set(RegistryContextKey, record)
			c.Next()
		}
	}
}

// MustRegistries validates the registries access.
func MustRegistries(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		registry := Current(c)

		if registry == nil {
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
			case action == "display": // && registry.Permission.DisplayRegistries:
				c.Next()
			case action == "change": // && registry.Permission.ChangeRegistries:
				c.Next()
			case action == "delete": // && registry.Permission.DeleteRegistries:
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
