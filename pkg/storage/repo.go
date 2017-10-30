package storage

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetRepos retrieves all available repos from storage.
func (db *data) GetRepos(filter *model.ReposFilter) (*model.Repos, error) {
	records := &model.Repos{}

	err := db.engine.Order(
		"name ASC",
	).Preload(
		"Org",
	).Preload(
		"Org.Registry",
	).Preload(
		"Tags",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateRepo creates an new repo.
func (db *data) CreateRepo(record *model.Repo) error {
	return db.engine.Create(
		&record,
	).Error
}

// UpdateRepo updates an repo.
func (db *data) UpdateRepo(record *model.Repo) error {
	return db.engine.Save(
		&record,
	).Error
}

// DeleteRepo deletes an repo.
func (db *data) DeleteRepo(record *model.Repo) error {
	return db.engine.Delete(
		&record,
	).Error
}

// GetRepo retrieves a specific repo from storage.
func (db *data) GetRepo(id string) (*model.Repo, error) {
	var (
		record = &model.Repo{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.engine.Where(
			"id = ?",
			val,
		)
	} else {
		query = db.engine.Where(
			"slug = ?",
			id,
		)
	}

	err := query.Model(
		&record,
	).Preload(
		"Org",
	).Preload(
		"Org.Registry",
	).Preload(
		"Tags",
	).First(
		&record,
	).Error

	return record, err
}
