package data

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetOrgs retrieves all available orgs from the database.
func (db *data) GetOrgs() (*model.Orgs, error) {
	records := &model.Orgs{}

	err := db.Order(
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

// CreateOrg creates a new org.
func (db *data) CreateOrg(record *model.Org, current *model.User) error {
	record.UserOrgs = model.UserOrgs{
		&model.UserOrg{
			UserID: current.ID,
			Perm:   "owner",
		},
	}

	return db.Create(
		&record,
	).Error
}

// UpdateOrg updates a org.
func (db *data) UpdateOrg(record *model.Org, current *model.User) error {
	return db.Save(
		&record,
	).Error
}

// DeleteOrg deletes a org.
func (db *data) DeleteOrg(record *model.Org, current *model.User) error {
	return db.Delete(
		&record,
	).Error
}

// GetOrg retrieves a specific org from the database.
func (db *data) GetOrg(id string) (*model.Org, *gorm.DB) {
	var (
		record = &model.Org{}
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
		"Registry",
	).Preload(
		"Repos",
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
func (db *data) GetOrgUsers(params *model.OrgUserParams) (*model.UserOrgs, error) {
	org, _ := db.GetOrg(params.Org)
	records := &model.UserOrgs{}

	err := db.Where(
		"org_id = ?",
		org.ID,
	).Model(
		&model.UserOrg{},
	).Preload(
		"User",
	).Preload(
		"Org",
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

func (db *data) CreateOrgUser(params *model.OrgUserParams, current *model.User) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.UserOrg{
					OrgID:  org.ID,
					UserID: user.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateOrgUser(params *model.OrgUserParams, current *model.User) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	return db.Model(
		&model.UserOrg{},
	).Where(
		"org_id = ? AND user_id = ?",
		org.ID,
		user.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteOrgUser(params *model.OrgUserParams, current *model.User) error {
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
func (db *data) GetOrgTeams(params *model.OrgTeamParams) (*model.TeamOrgs, error) {
	org, _ := db.GetOrg(params.Org)
	records := &model.TeamOrgs{}

	err := db.Where(
		"org_id = ?",
		org.ID,
	).Model(
		&model.TeamOrg{},
	).Preload(
		"Team",
	).Preload(
		"Org",
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

func (db *data) CreateOrgTeam(params *model.OrgTeamParams, current *model.User) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamOrg{
					OrgID:  org.ID,
					TeamID: team.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateOrgTeam(params *model.OrgTeamParams, current *model.User) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		&model.TeamOrg{},
	).Where(
		"org_id = ? AND team_id = ?",
		org.ID,
		team.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteOrgTeam(params *model.OrgTeamParams, current *model.User) error {
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
