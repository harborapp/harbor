package storage

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetUsers retrieves all available users from storage.
func (db *data) GetUsers() (*model.Users, error) {
	records := &model.Users{}

	err := db.engine.Order(
		"username ASC",
	).Preload(
		"Teams",
	).Preload(
		"Orgs",
	).Find(
		records,
	).Error

	return records, err
}

// CreateUser creates an new user.
func (db *data) CreateUser(record *model.User) error {
	return db.engine.Create(
		record,
	).Error
}

// UpdateUser updates an user.
func (db *data) UpdateUser(record *model.User) error {
	return db.engine.Save(
		record,
	).Error
}

// DeleteUser deletes an user.
func (db *data) DeleteUser(record *model.User) error {
	return db.engine.Delete(
		record,
	).Error
}

// GetUser retrieves a specific user from storage.
func (db *data) GetUser(id string) (*model.User, error) {
	var (
		record = &model.User{}
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
		"Teams",
	).Preload(
		"Orgs",
	).First(
		record,
	).Error

	return record, err
}
