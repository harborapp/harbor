package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/pkg/version"
)

// IndexInfo represents the API index.
func IndexInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"api":     "Umschlag API",
			"version": version.Version.String(),
			"stream":  "master",
		},
	)
}
