package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
	"golang.org/x/net/context"
)

// GetRegistries retrieves all available registries from the database.
func GetRegistries(c context.Context) (*model.Registries, error) {
	return FromContext(c).GetRegistries()
}

// CreateRegistry creates a new registry.
func CreateRegistry(c context.Context, record *model.Registry) error {
	return FromContext(c).CreateRegistry(record, Current(c))
}

// UpdateRegistry updates a registry.
func UpdateRegistry(c context.Context, record *model.Registry) error {
	return FromContext(c).UpdateRegistry(record, Current(c))
}

// DeleteRegistry deletes a registry.
func DeleteRegistry(c context.Context, record *model.Registry) error {
	return FromContext(c).DeleteRegistry(record, Current(c))
}

// GetRegistry retrieves a specific registry from the database.
func GetRegistry(c context.Context, id string) (*model.Registry, *gorm.DB) {
	return FromContext(c).GetRegistry(id)
}
