package storage

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetRegistries retrieves all available registries from storage.
func (db *data) GetRegistries() (*model.Registries, error) {
	records := &model.Registries{}

	err := db.engine.Order(
		"name ASC",
	).Preload(
		"Orgs",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateRegistry creates an new registry.
func (db *data) CreateRegistry(record *model.Registry) error {
	return db.engine.Create(
		&record,
	).Error
}

// UpdateRegistry updates an registry.
func (db *data) UpdateRegistry(record *model.Registry) error {
	return db.engine.Save(
		&record,
	).Error
}

// DeleteRegistry deletes an registry.
func (db *data) DeleteRegistry(record *model.Registry) error {
	return db.engine.Delete(
		&record,
	).Error
}

// GetRegistry retrieves a specific registry from storage.
func (db *data) GetRegistry(id string) (*model.Registry, error) {
	var (
		record = &model.Registry{}
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
		"Orgs",
	).First(
		&record,
	).Error

	return record, err
}
