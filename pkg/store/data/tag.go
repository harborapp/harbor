package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetTags retrieves all available tags from the database.
func (db *data) GetTags() (*model.Tags, error) {
	records := &model.Tags{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Repo",
	).Preload(
		"Repo.Org",
	).Preload(
		"Repo.Org.Registry",
	).Find(
		records,
	).Error

	return records, err
}

// CreateTag creates a new tag.
func (db *data) CreateTag(record *model.Tag, current *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateTag updates a tag.
func (db *data) UpdateTag(record *model.Tag, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteTag deletes a tag.
func (db *data) DeleteTag(record *model.Tag, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetTag retrieves a specific tag from the database.
func (db *data) GetTag(id string) (*model.Tag, *gorm.DB) {
	var (
		record = &model.Tag{}
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
		record,
	).Preload(
		"Repo",
	).Preload(
		"Repo.Org",
	).Preload(
		"Repo.Org.Registry",
	).First(
		record,
	)

	return record, res
}
