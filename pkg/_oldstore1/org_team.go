package storage

import (
	"errors"

	"github.com/umschlag/umschlag-api/pkg/model"
)

var (
	// ErrInvalidOrgTeamPerm defines the error for invalid permissions.
	ErrInvalidOrgTeamPerm = errors.New("invalid permission, can be user, admin or owner")
)

// GetOrgTeams retrieves teams for an org.
func (db *data) GetOrgTeams(params *model.OrgTeamParams) (*model.TeamOrgs, error) {
	org, _ := db.GetOrg(params.Org)
	records := &model.TeamOrgs{}

	err := db.engine.Where(
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

// GetOrgHasTeam checks if an specific team is assigned to an org.
func (db *data) GetOrgHasTeam(params *model.OrgTeamParams) bool {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	res := db.engine.Model(
		org,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

// CreateOrgTeam assigns a team to an specific org.
func (db *data) CreateOrgTeam(params *model.OrgTeamParams) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.engine.Create(
				&model.TeamOrg{
					OrgID:  org.ID,
					TeamID: team.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return ErrInvalidOrgTeamPerm
}

// UpdateOrgTeam updates the org team permission.
func (db *data) UpdateOrgTeam(params *model.OrgTeamParams) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	return db.engine.Model(
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

// DeleteOrgTeam removes a team from an specific org.
func (db *data) DeleteOrgTeam(params *model.OrgTeamParams) error {
	org, _ := db.GetOrg(params.Org)
	team, _ := db.GetTeam(params.Team)

	return db.engine.Model(
		org,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
