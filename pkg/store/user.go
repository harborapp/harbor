package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
	"golang.org/x/net/context"
)

// GetUsers retrieves all available users from the database.
func GetUsers(c context.Context) (*model.Users, error) {
	return FromContext(c).GetUsers()
}

// CreateUser creates a new user.
func CreateUser(c context.Context, record *model.User) error {
	return FromContext(c).CreateUser(record, Current(c))
}

// UpdateUser updates a user.
func UpdateUser(c context.Context, record *model.User) error {
	return FromContext(c).UpdateUser(record, Current(c))
}

// DeleteUser deletes a user.
func DeleteUser(c context.Context, record *model.User) error {
	return FromContext(c).DeleteUser(record, Current(c))
}

// GetUser retrieves a specific user from the database.
func GetUser(c context.Context, id string) (*model.User, *gorm.DB) {
	return FromContext(c).GetUser(id)
}

// GetUserTeams retrieves teams for a user.
func GetUserTeams(c context.Context, params *model.UserTeamParams) (*model.TeamUsers, error) {
	return FromContext(c).GetUserTeams(params)
}

// GetUserHasTeam checks if a specific team is assigned to a user.
func GetUserHasTeam(c context.Context, params *model.UserTeamParams) bool {
	return FromContext(c).GetUserHasTeam(params)
}

// CreateUserTeam assigns a team to a specific user.
func CreateUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).CreateUserTeam(params, Current(c))
}

// UpdateUserTeam updates the user team permission.
func UpdateUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).UpdateUserTeam(params, Current(c))
}

// DeleteUserTeam removes a team from a specific user.
func DeleteUserTeam(c context.Context, params *model.UserTeamParams) error {
	return FromContext(c).DeleteUserTeam(params, Current(c))
}

// GetUserOrgs retrieves orgs for a user.
func GetUserOrgs(c context.Context, params *model.UserOrgParams) (*model.UserOrgs, error) {
	return FromContext(c).GetUserOrgs(params)
}

// GetUserHasOrg checks if a specific org is assigned to a user.
func GetUserHasOrg(c context.Context, params *model.UserOrgParams) bool {
	return FromContext(c).GetUserHasOrg(params)
}

// CreateUserOrg assigns a org to a specific user.
func CreateUserOrg(c context.Context, params *model.UserOrgParams) error {
	return FromContext(c).CreateUserOrg(params, Current(c))
}

// UpdateUserOrg updates the user org permission.
func UpdateUserOrg(c context.Context, params *model.UserOrgParams) error {
	return FromContext(c).UpdateUserOrg(params, Current(c))
}

// DeleteUserOrg removes a org from a specific user.
func DeleteUserOrg(c context.Context, params *model.UserOrgParams) error {
	return FromContext(c).DeleteUserOrg(params, Current(c))
}
