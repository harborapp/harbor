package data

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetOrgs retrieves all available orgs from the database.
func (db *data) GetOrgs() (*model.Orgs, error) {
	records := &model.Orgs{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Registry",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateOrg creates a new org.
func (db *data) CreateOrg(record *model.Org) error {
	return db.Create(
		&record,
	).Error
}

// UpdateOrg updates a org.
func (db *data) UpdateOrg(record *model.Org) error {
	return db.Save(
		&record,
	).Error
}

// DeleteOrg deletes a org.
func (db *data) DeleteOrg(record *model.Org) error {
	return db.Delete(
		&record,
	).Error
}

// GetOrg retrieves a specific org from the database.
func (db *data) GetOrg(id string) (*model.Org, *gorm.DB) {
	record := &model.Org{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Registry",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).First(
		&record,
	)

	return record, res
}

// GetOrgUsers retrieves users for a org.
func (db *data) GetOrgUsers(params *model.OrgUserParams) (*model.Users, error) {
	org, _ := db.GetOrg(params.Org)

	records := &model.Users{}

	err := db.Model(
		org,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetOrgHasUser checks if a specific user is assigned to a org.
func (db *data) GetOrgHasUser(params *model.OrgUserParams) bool {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	res := db.Model(
		org,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreateOrgUser(params *model.OrgUserParams) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	return db.Model(
		org,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeleteOrgUser(params *model.OrgUserParams) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	return db.Model(
		org,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetOrgTeams retrieves teams for a org.
func (db *data) GetOrgTeams(params *model.OrgTeamParams) (*model.Teams, error) {
	org, _ := db.GetOrg(params.Org)

	records := &model.Teams{}

	err := db.Model(
		org,
	).Association(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// GetOrgHasTeam checks if a specific team is assigned to a org.
func (db *data) GetOrgHasTeam(params *model.OrgTeamParams) bool {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	res := db.Model(
		org,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

func (db *data) CreateOrgTeam(params *model.OrgTeamParams) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		org,
	).Association(
		"Teams",
	).Append(
		team,
	).Error
}

func (db *data) DeleteOrgTeam(params *model.OrgTeamParams) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		org,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
