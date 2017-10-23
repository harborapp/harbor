package data

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetUsers retrieves all available users from the database.
func (db *data) GetUsers() (*model.Users, error) {
	records := &model.Users{}

	err := db.Order(
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

// CreateUser creates a new user.
func (db *data) CreateUser(record *model.User, current *model.User) error {
	return db.Create(
		record,
	).Error
}

// UpdateUser updates a user.
func (db *data) UpdateUser(record *model.User, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteUser deletes a user.
func (db *data) DeleteUser(record *model.User, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetUser retrieves a specific user from the database.
func (db *data) GetUser(id string) (*model.User, *gorm.DB) {
	var (
		record = &model.User{}
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
		record,
	).Preload(
		"Teams",
	).Preload(
		"Orgs",
	).First(
		record,
	)

	return record, res
}

// GetUserTeams retrieves teams for a user.
func (db *data) GetUserTeams(params *model.UserTeamParams) (*model.TeamUsers, error) {
	user, _ := db.GetUser(params.User)
	records := &model.TeamUsers{}

	err := db.Where(
		"user_id = ?",
		user.ID,
	).Model(
		&model.TeamUser{},
	).Preload(
		"User",
	).Preload(
		"Team",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasTeam checks if a specific team is assigned to a user.
func (db *data) GetUserHasTeam(params *model.UserTeamParams) bool {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	res := db.Model(
		user,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

func (db *data) CreateUserTeam(params *model.UserTeamParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamUser{
					UserID: user.ID,
					TeamID: team.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateUserTeam(params *model.UserTeamParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		&model.TeamUser{},
	).Where(
		"user_id = ? AND team_id = ?",
		user.ID,
		team.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteUserTeam(params *model.UserTeamParams, current *model.User) error {
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
func (db *data) GetUserOrgs(params *model.UserOrgParams) (*model.UserOrgs, error) {
	user, _ := db.GetUser(params.User)
	records := &model.UserOrgs{}

	err := db.Where(
		"user_id = ?",
		user.ID,
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

// GetUserHasOrg checks if a specific org is assigned to a user.
func (db *data) GetUserHasOrg(params *model.UserOrgParams) bool {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	res := db.Model(
		user,
	).Association(
		"Orgs",
	).Find(
		org,
	).Error

	return res == nil
}

func (db *data) CreateUserOrg(params *model.UserOrgParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.UserOrg{
					UserID: user.ID,
					OrgID:  org.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateUserOrg(params *model.UserOrgParams, current *model.User) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	return db.Model(
		&model.UserOrg{},
	).Where(
		"user_id = ? AND org_id = ?",
		user.ID,
		org.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteUserOrg(params *model.UserOrgParams, current *model.User) error {
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
