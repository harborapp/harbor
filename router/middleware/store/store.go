package store

import (
	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/store"
	"github.com/umschlag/umschlag-api/store/data"
)

// SetStore injects the storage into the context.
func SetStore() gin.HandlerFunc {
	db := data.Load()

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
