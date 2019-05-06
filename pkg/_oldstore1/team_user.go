package storage

import (
	"errors"

	"github.com/umschlag/umschlag-api/pkg/model"
)

var (
	// ErrInvalidTeamUserPerm defines the error for invalid permissions.
	ErrInvalidTeamUserPerm = errors.New("invalid permission, can be user, admin or owner")
)

// GetTeamUsers retrieves users for a team.
func (db *data) GetTeamUsers(params *model.TeamUserParams) (*model.TeamUsers, error) {
	team, _ := db.GetTeam(params.Team)
	records := &model.TeamUsers{}

	err := db.engine.Where(
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

	res := db.engine.Model(
		team,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

func (db *data) CreateTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.engine.Create(
				&model.TeamUser{
					TeamID: team.ID,
					UserID: user.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return ErrInvalidTeamUserPerm
}

func (db *data) UpdateTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.engine.Model(
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

func (db *data) DeleteTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.engine.Model(
		team,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}
