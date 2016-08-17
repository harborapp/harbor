package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
	"golang.org/x/net/context"
)

// GetRepositories retrieves all available repositories from the database.
func GetRepositories(c context.Context) (*model.Repositories, error) {
	return FromContext(c).GetRepositories()
}

// CreateRepository creates a new repository.
func CreateRepository(c context.Context, record *model.Repository) error {
	return FromContext(c).CreateRepository(record)
}

// UpdateRepository updates a repository.
func UpdateRepository(c context.Context, record *model.Repository) error {
	return FromContext(c).UpdateRepository(record)
}

// DeleteRepository deletes a repository.
func DeleteRepository(c context.Context, record *model.Repository) error {
	return FromContext(c).DeleteRepository(record)
}

// GetRepository retrieves a specific repository from the database.
func GetRepository(c context.Context, id string) (*model.Repository, *gorm.DB) {
	return FromContext(c).GetRepository(id)
}
