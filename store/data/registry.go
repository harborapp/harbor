package data

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetRegistries retrieves all available registries from the database.
func (db *data) GetRegistries() (*model.Registries, error) {
	records := &model.Registries{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Orgs",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateRegistry creates a new registry.
func (db *data) CreateRegistry(record *model.Registry) error {
	return db.Create(
		&record,
	).Error
}

// UpdateRegistry updates a registry.
func (db *data) UpdateRegistry(record *model.Registry) error {
	return db.Save(
		&record,
	).Error
}

// DeleteRegistry deletes a registry.
func (db *data) DeleteRegistry(record *model.Registry) error {
	return db.Delete(
		&record,
	).Error
}

// GetRegistry retrieves a specific registry from the database.
func (db *data) GetRegistry(id string) (*model.Registry, *gorm.DB) {
	var (
		record = &model.Registry{}
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
		"Orgs",
	).First(
		&record,
	)

	return record, res
}
