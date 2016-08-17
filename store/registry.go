package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
	"golang.org/x/net/context"
)

// GetRegistries retrieves all available registries from the database.
func GetRegistries(c context.Context) (*model.Registries, error) {
	return FromContext(c).GetRegistries()
}

// CreateRegistry creates a new registry.
func CreateRegistry(c context.Context, record *model.Registry) error {
	return FromContext(c).CreateRegistry(record)
}

// UpdateRegistry updates a registry.
func UpdateRegistry(c context.Context, record *model.Registry) error {
	return FromContext(c).UpdateRegistry(record)
}

// DeleteRegistry deletes a registry.
func DeleteRegistry(c context.Context, record *model.Registry) error {
	return FromContext(c).DeleteRegistry(record)
}

// GetRegistry retrieves a specific registry from the database.
func GetRegistry(c context.Context, id string) (*model.Registry, *gorm.DB) {
	return FromContext(c).GetRegistry(id)
}
