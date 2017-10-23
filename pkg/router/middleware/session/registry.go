package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/store"
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
		current := Current(c)

		if current.Admin {
			c.Next()
			return
		}

		switch {
		case action == "display":
			if allowRegistryDisplay(c) {
				c.Next()
				return
			}
		case action == "change":
			if allowRegistryChange(c) {
				c.Next()
				return
			}
		case action == "delete":
			if allowRegistryDelete(c) {
				c.Next()
				return
			}
		}

		AbortUnauthorized(c)
	}
}

// allowRegistryDisplay checks if the given user is allowed to display the resource.
func allowRegistryDisplay(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowRegistryChange checks if the given user is allowed to change the resource.
func allowRegistryChange(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}

// allowRegistryDelete checks if the given user is allowed to delete the resource.
func allowRegistryDelete(c *gin.Context) bool {
	// TODO(tboerger): Add real implementation
	return false
}
