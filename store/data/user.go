package data

import (
	"github.com/harborapp/harbor-api/model"
	"github.com/jinzhu/gorm"
)

// GetUsers retrieves all available users from the database.
func (db *data) GetUsers() (*model.Users, error) {
	records := &model.Users{}

	err := db.Order(
		"username ASC",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateUser creates a new user.
func (db *data) CreateUser(record *model.User) error {
	return db.Create(
		&record,
	).Error
}

// UpdateUser updates a user.
func (db *data) UpdateUser(record *model.User) error {
	return db.Save(
		&record,
	).Error
}

// DeleteUser deletes a user.
func (db *data) DeleteUser(record *model.User) error {
	return db.Delete(
		&record,
	).Error
}

// GetUser retrieves a specific user from the database.
func (db *data) GetUser(id string) (*model.User, *gorm.DB) {
	record := &model.User{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).First(
		&record,
	)

	return record, res
}
