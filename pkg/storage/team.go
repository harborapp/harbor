package storage

import (
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetTeams retrieves all available teams from storage.
func (db *data) GetTeams() (*model.Teams, error) {
	records := &model.Teams{}

	err := db.engine.Order(
		"name ASC",
	).Preload(
		"Users",
	).Preload(
		"Orgs",
	).Find(
		records,
	).Error

	return records, err
}

// CreateTeam creates an new team.
func (db *data) CreateTeam(record *model.Team) error {

	// TODO(tboerger): inject current user

	// record.TeamUsers = model.TeamUsers{
	// 	&model.TeamUser{
	// 		UserID: current.ID,
	// 		Perm:   "owner",
	// 	},
	// }

	return db.engine.Create(
		record,
	).Error
}

// UpdateTeam updates an team.
func (db *data) UpdateTeam(record *model.Team) error {
	return db.engine.Save(
		record,
	).Error
}

// DeleteTeam deletes an team.
func (db *data) DeleteTeam(record *model.Team) error {
	return db.engine.Delete(
		record,
	).Error
}

// GetTeam retrieves a specific team from storage.
func (db *data) GetTeam(id string) (*model.Team, error) {
	var (
		record = &model.Team{}
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
		"Users",
	).Preload(
		"Orgs",
	).First(
		record,
	).Error

	return record, err
}
