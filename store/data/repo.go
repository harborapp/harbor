package data

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetRepos retrieves all available repos from the database.
func (db *data) GetRepos() (*model.Repos, error) {
	records := &model.Repos{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Tags",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateRepo creates a new repo.
func (db *data) CreateRepo(record *model.Repo) error {
	return db.Create(
		&record,
	).Error
}

// UpdateRepo updates a repo.
func (db *data) UpdateRepo(record *model.Repo) error {
	return db.Save(
		&record,
	).Error
}

// DeleteRepo deletes a repo.
func (db *data) DeleteRepo(record *model.Repo) error {
	return db.Delete(
		&record,
	).Error
}

// GetRepo retrieves a specific repo from the database.
func (db *data) GetRepo(id string) (*model.Repo, *gorm.DB) {
	record := &model.Repo{}

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
