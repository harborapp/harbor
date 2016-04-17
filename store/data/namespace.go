package data

import (
	"github.com/harborapp/harbor-api/model"
	"github.com/jinzhu/gorm"
)

// GetNamespaces retrieves all available namespaces from the database.
func (db *data) GetNamespaces() (*model.Namespaces, error) {
	records := &model.Namespaces{}

	err := db.Order(
		"name ASC",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateNamespace creates a new namespace.
func (db *data) CreateNamespace(record *model.Namespace) error {
	return db.Create(
		&record,
	).Error
}

// UpdateNamespace updates a namespace.
func (db *data) UpdateNamespace(record *model.Namespace) error {
	return db.Save(
		&record,
	).Error
}

// DeleteNamespace deletes a namespace.
func (db *data) DeleteNamespace(record *model.Namespace) error {
	return db.Delete(
		&record,
	).Error
}

// GetNamespace retrieves a specific namespace from the database.
func (db *data) GetNamespace(id string) (*model.Namespace, *gorm.DB) {
	record := &model.Namespace{}

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
