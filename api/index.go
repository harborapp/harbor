package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harborapp/harbor-api/config"
)

// IndexInfo represents the API index.
func IndexInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"api":     "Harbor API",
			"version": config.Version,
			"stream":  "master",
		},
	)
}
