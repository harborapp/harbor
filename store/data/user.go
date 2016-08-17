package data

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetUsers retrieves all available users from the database.
func (db *data) GetUsers() (*model.Users, error) {
	records := &model.Users{}

	err := db.Order(
		"username ASC",
	).Find(
		records,
	).Error

	return records, err
}

// CreateUser creates a new user.
func (db *data) CreateUser(record *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateUser updates a user.
func (db *data) UpdateUser(record *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteUser deletes a user.
func (db *data) DeleteUser(record *model.User) error {
	return db.Delete(
		record,
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
		record,
	).First(
		record,
	)

	return record, res
}

// GetUserTeams retrieves teams for a user.
func (db *data) GetUserTeams(params *model.UserTeamParams) (*model.Teams, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Teams{}

	err := db.Model(
		user,
	).Association(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasTeam checks if a specific team is assigned to a user.
func (db *data) GetUserHasTeam(params *model.UserTeamParams) bool {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	count := db.Model(
		user,
	).Association(
		"Teams",
	).Find(
		team,
	).Count()

	return count > 0
}

func (db *data) CreateUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		user,
	).Association(
		"Teams",
	).Append(
		team,
	).Error
}

func (db *data) DeleteUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		user,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}

// GetUserOrgs retrieves orgs for a user.
func (db *data) GetUserOrgs(params *model.UserOrgParams) (*model.Orgs, error) {
	user, _ := db.GetUser(params.User)

	records := &model.Orgs{}

	err := db.Model(
		user,
	).Association(
		"Orgs",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasOrg checks if a specific org is assigned to a user.
func (db *data) GetUserHasOrg(params *model.UserOrgParams) bool {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	count := db.Model(
		user,
	).Association(
		"Orgs",
	).Find(
		org,
	).Count()

	return count > 0
}

func (db *data) CreateUserOrg(params *model.UserOrgParams) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	return db.Model(
		user,
	).Association(
		"Orgs",
	).Append(
		org,
	).Error
}

func (db *data) DeleteUserOrg(params *model.UserOrgParams) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	return db.Model(
		user,
	).Association(
		"Orgs",
	).Delete(
		org,
	).Error
}
