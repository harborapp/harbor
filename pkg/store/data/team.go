package data

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
)

// GetTeams retrieves all available teams from the database.
func (db *data) GetTeams() (*model.Teams, error) {
	records := &model.Teams{}

	err := db.Order(
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

// CreateTeam creates a new team.
func (db *data) CreateTeam(record *model.Team, current *model.User) error {
	record.TeamUsers = model.TeamUsers{
		&model.TeamUser{
			UserID: current.ID,
			Perm:   "owner",
		},
	}

	return db.Create(
		record,
	).Error
}

// UpdateTeam updates a team.
func (db *data) UpdateTeam(record *model.Team, current *model.User) error {
	return db.Save(
		record,
	).Error
}

// DeleteTeam deletes a team.
func (db *data) DeleteTeam(record *model.Team, current *model.User) error {
	return db.Delete(
		record,
	).Error
}

// GetTeam retrieves a specific team from the database.
func (db *data) GetTeam(id string) (*model.Team, *gorm.DB) {
	var (
		record = &model.Team{}
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
		"Users",
	).Preload(
		"Orgs",
	).First(
		record,
	)

	return record, res
}

// GetTeamUsers retrieves users for a team.
func (db *data) GetTeamUsers(params *model.TeamUserParams) (*model.TeamUsers, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamUsers{}

	err := db.Where(
		"team_id = ?",
		team.ID,
	).Model(
		&model.TeamUser{},
	).Preload(
		"Team",
	).Preload(
		"User",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasUser checks if a specific user is assigned to a team.
func (db *data) GetTeamHasUser(params *model.TeamUserParams) bool {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	res := db.Model(
		team,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreateTeamUser(params *model.TeamUserParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamUser{
					TeamID: team.ID,
					UserID: user.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateTeamUser(params *model.TeamUserParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		&model.TeamUser{},
	).Where(
		"team_id = ? AND user_id = ?",
		team.ID,
		user.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteTeamUser(params *model.TeamUserParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		team,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetTeamOrgs retrieves orgs for a team.
func (db *data) GetTeamOrgs(params *model.TeamOrgParams) (*model.TeamOrgs, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamOrgs{}

	err := db.Where(
		"team_id = ?",
		team.ID,
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

// GetTeamHasOrg checks if a specific org is assigned to a team.
func (db *data) GetTeamHasOrg(params *model.TeamOrgParams) bool {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	res := db.Model(
		team,
	).Association(
		"Orgs",
	).Find(
		org,
	).Error

	return res == nil
}

func (db *data) CreateTeamOrg(params *model.TeamOrgParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.Create(
				&model.TeamOrg{
					TeamID: team.ID,
					OrgID:  org.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return fmt.Errorf("Invalid permission, can be user, admin or owner")
}

func (db *data) UpdateTeamOrg(params *model.TeamOrgParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	return db.Model(
		&model.TeamOrg{},
	).Where(
		"team_id = ? AND org_id = ?",
		team.ID,
		org.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

func (db *data) DeleteTeamOrg(params *model.TeamOrgParams, current *model.User) error {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	return db.Model(
		team,
	).Association(
		"Orgs",
	).Delete(
		org,
	).Error
}
