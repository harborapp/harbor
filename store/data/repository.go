package data

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetRepositories retrieves all available repositories from the database.
func (db *data) GetRepositories() (*model.Repositories, error) {
	records := &model.Repositories{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Tags",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateRepository creates a new repository.
func (db *data) CreateRepository(record *model.Repository) error {
	return db.Create(
		&record,
	).Error
}

// UpdateRepository updates a repository.
func (db *data) UpdateRepository(record *model.Repository) error {
	return db.Save(
		&record,
	).Error
}

// DeleteRepository deletes a repository.
func (db *data) DeleteRepository(record *model.Repository) error {
	return db.Delete(
		&record,
	).Error
}

// GetRepository retrieves a specific repository from the database.
func (db *data) GetRepository(id string) (*model.Repository, *gorm.DB) {
	record := &model.Repository{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Tags",
	).First(
		&record,
	)

	return record, res
}
