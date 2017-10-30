package storage

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetTags retrieves all available tags from storage.
func (db *data) GetTags(tags *model.TagsFilter) (*model.Tags, error) {
	records := &model.Tags{}

	err := db.engine.Order(
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

// CreateTag creates an new tag.
func (db *data) CreateTag(record *model.Tag) error {
	return db.engine.Create(
		record,
	).Error
}

// UpdateTag updates an tag.
func (db *data) UpdateTag(record *model.Tag) error {
	return db.engine.Save(
		record,
	).Error
}

// DeleteTag deletes an tag.
func (db *data) DeleteTag(record *model.Tag) error {
	return db.engine.Delete(
		record,
	).Error
}

// GetTag retrieves a specific tag from storage.
func (db *data) GetTag(id string) (*model.Tag, error) {
	var (
		record = &model.Tag{}
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
		record,
	).Preload(
		"Repo",
	).Preload(
		"Repo.Org",
	).Preload(
		"Repo.Org.Registry",
	).First(
		record,
	).Error

	return record, err
}
