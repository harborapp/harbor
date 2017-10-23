package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/umschlag/umschlag-api/pkg/model"
	"github.com/umschlag/umschlag-api/pkg/router/middleware/session"
	"github.com/umschlag/umschlag-api/pkg/store"
)

// RegistryIndex retrieves all available registries.
func RegistryIndex(c *gin.Context) {
	records, err := store.GetRegistries(
		c,
	)

	if err != nil {
		logrus.Warnf("Failed to fetch registries. %s", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch registries",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// RegistryShow retrieves a specific registry.
func RegistryShow(c *gin.Context) {
	record := session.Registry(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// RegistryDelete removes a specific registry.
func RegistryDelete(c *gin.Context) {
	record := session.Registry(c)

	err := store.DeleteRegistry(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to delete registry. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to delete registry",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted registry",
		},
	)
}

// RegistryUpdate updates an existing registry.
func RegistryUpdate(c *gin.Context) {
	record := session.Registry(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind registry data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind registry data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateRegistry(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to update registry. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// RegistryCreate creates a new user.
func RegistryCreate(c *gin.Context) {
	record := &model.Registry{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warnf("Failed to bind registry data. %s", err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind registry data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateRegistry(
		c,
		record,
	)

	if err != nil {
		logrus.Warnf("Failed to create registry. %s", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}
