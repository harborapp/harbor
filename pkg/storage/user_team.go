package storage

import (
	"errors"

	"github.com/umschlag/umschlag-api/pkg/model"
)

var (
	// ErrInvalidUserTeamPerm defines the error for invalid permissions.
	ErrInvalidUserTeamPerm = errors.New("invalid permission, can be user, admin or owner")
)

// GetUserTeams retrieves teams for an user.
func (db *data) GetUserTeams(params *model.UserTeamParams) (*model.TeamUsers, error) {
	user, _ := db.GetUser(params.User)
	records := &model.TeamUsers{}

	err := db.engine.Where(
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

// GetUserHasTeam checks if a specific team is assigned to an user.
func (db *data) GetUserHasTeam(params *model.UserTeamParams) bool {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	res := db.engine.Model(
		user,
	).Association(
		"Teams",
	).Find(
		team,
	).Error

	return res == nil
}

// CreateUserTeam assigns a team to an specific user.
func (db *data) CreateUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.engine.Create(
				&model.TeamUser{
					UserID: user.ID,
					TeamID: team.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return ErrInvalidUserTeamPerm
}

// UpdateUserTeam updates the user team permission.
func (db *data) UpdateUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.engine.Model(
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

// DeleteUserTeam removes a team from an specific user.
func (db *data) DeleteUserTeam(params *model.UserTeamParams) error {
	user, _ := db.GetUser(params.User)
	team, _ := db.GetTeam(params.Team)

	return db.engine.Model(
		user,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
