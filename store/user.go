package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
	"golang.org/x/net/context"
)

// GetUsers retrieves all available users from the database.
func GetUsers(c context.Context) (*model.Users, error) {
	return FromContext(c).GetUsers()
}

// CreateUser creates a new user.
func CreateUser(c context.Context, record *model.User) error {
	return FromContext(c).CreateUser(record)
}

// UpdateUser updates a user.
func UpdateUser(c context.Context, record *model.User) error {
	return FromContext(c).UpdateUser(record)
}

// DeleteUser deletes a user.
func DeleteUser(c context.Context, record *model.User) error {
	return FromContext(c).DeleteUser(record)
}

// GetUser retrieves a specific user from the database.
func GetUser(c context.Context, id string) (*model.User, *gorm.DB) {
	return FromContext(c).GetUser(id)
}

// GetUserTeams retrieves teams for a user.
func GetUserTeams(c context.Context, params *model.UserTeamParams) (*model.Teams, error) {
	return FromContext(c).GetUserTeams(params)
}

// GetUserHasTeam checks if a specific team is assigned to a user.
func GetUserHasTeam(c context.Context, params *model.UserTeamParams) bool {
	return FromContext(c).GetUserHasTeam(params)
}

// CreateUserTeam assigns a team to a specific user.
func CreateUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).CreateUserTeam(params)
}

// DeleteUserTeam removes a team from a specific user.
func DeleteUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).DeleteUserTeam(params)
}
