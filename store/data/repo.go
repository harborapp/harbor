package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetRepos retrieves all available repos from the database.
func (db *data) GetRepos() (*model.Repos, error) {
	records := &model.Repos{}

	err := db.Order(
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

// CreateRepo creates a new repo.
func (db *data) CreateRepo(record *model.Repo, current *model.User) error {
	return db.Create(
		&record,
	).Error
}

// UpdateRepo updates a repo.
func (db *data) UpdateRepo(record *model.Repo, current *model.User) error {
	return db.Save(
		&record,
	).Error
}

// DeleteRepo deletes a repo.
func (db *data) DeleteRepo(record *model.Repo, current *model.User) error {
	return db.Delete(
		&record,
	).Error
}

// GetRepo retrieves a specific repo from the database.
func (db *data) GetRepo(id string) (*model.Repo, *gorm.DB) {
	var (
		record = &model.Repo{}
		query  *gorm.DB
	)

	if match, _ := regexp.MatchString("^([0-9]+)$", id); match {
		val, _ := strconv.ParseInt(id, 10, 64)

		query = db.Where(
			"id = ?",
			val,
		)
	} else {
		query = db.Where(
			"slug = ?",
			id,
		)
	}

	res := query.Model(
		&record,
	).Preload(
		"Org",
	).Preload(
		"Org.Registry",
	).Preload(
		"Tags",
	).First(
		&record,
	)

	return record, res
}
