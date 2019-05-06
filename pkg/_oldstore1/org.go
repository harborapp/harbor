package storage

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetOrgs retrieves all available orgs from storage.
func (db *data) GetOrgs() (*model.Orgs, error) {
	records := &model.Orgs{}

	err := db.engine.Order(
		"name ASC",
	).Preload(
		"Registry",
	).Preload(
		"Repos",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateOrg creates an new org.
func (db *data) CreateOrg(record *model.Org) error {

	// TODO(tboerger): inject current user

	// record.UserOrgs = model.UserOrgs{
	// 	&model.UserOrg{
	// 		UserID: current.ID,
	// 		Perm:   "owner",
	// 	},
	// }

	return db.engine.Create(
		&record,
	).Error
}

// UpdateOrg updates an org.
func (db *data) UpdateOrg(record *model.Org) error {
	return db.engine.Save(
		&record,
	).Error
}

// DeleteOrg deletes an org.
func (db *data) DeleteOrg(record *model.Org) error {
	return db.engine.Delete(
		&record,
	).Error
}

// GetOrg retrieves a specific org from storage.
func (db *data) GetOrg(id string) (*model.Org, error) {
	var (
		record = &model.Org{}
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
		"Registry",
	).Preload(
		"Repos",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).First(
		&record,
	).Error

	return record, err
}
