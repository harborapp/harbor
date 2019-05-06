package storage

import (
	"errors"

	"github.com/umschlag/umschlag-api/pkg/model"
)

var (
	// ErrInvalidTeamOrgPerm defines the error for invalid permissions.
	ErrInvalidTeamOrgPerm = errors.New("invalid permission, can be user, admin or owner")
)

// GetTeamOrgs retrieves orgs for a team.
func (db *data) GetTeamOrgs(params *model.TeamOrgParams) (*model.TeamOrgs, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamOrgs{}

	err := db.engine.Where(
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

	res := db.engine.Model(
		team,
	).Association(
		"Orgs",
	).Find(
		org,
	).Error

	return res == nil
}

func (db *data) CreateTeamOrg(params *model.TeamOrgParams) error {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.engine.Create(
				&model.TeamOrg{
					TeamID: team.ID,
					OrgID:  org.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return ErrInvalidTeamOrgPerm
}

func (db *data) UpdateTeamOrg(params *model.TeamOrgParams) error {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	return db.engine.Model(
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

func (db *data) DeleteTeamOrg(params *model.TeamOrgParams) error {
	team, _ := db.GetTeam(params.Team)
	org, _ := db.GetOrg(params.Org)

	return db.engine.Model(
		team,
	).Association(
		"Orgs",
	).Delete(
		org,
	).Error
}
