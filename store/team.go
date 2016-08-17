package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
	"golang.org/x/net/context"
)

// GetTeams retrieves all available teams from the database.
func GetTeams(c context.Context) (*model.Teams, error) {
	return FromContext(c).GetTeams()
}

// CreateTeam creates a new team.
func CreateTeam(c context.Context, record *model.Team) error {
	return FromContext(c).CreateTeam(record)
}

// UpdateTeam updates a team.
func UpdateTeam(c context.Context, record *model.Team) error {
	return FromContext(c).UpdateTeam(record)
}

// DeleteTeam deletes a team.
func DeleteTeam(c context.Context, record *model.Team) error {
	return FromContext(c).DeleteTeam(record)
}

// GetTeam retrieves a specific team from the database.
func GetTeam(c context.Context, id string) (*model.Team, *gorm.DB) {
	return FromContext(c).GetTeam(id)
}

// GetTeamUsers retrieves users for a team.
func GetTeamUsers(c context.Context, params *model.TeamUserParams) (*model.Users, error) {
	return FromContext(c).GetTeamUsers(params)
}

// GetTeamHasUser checks if a specific user is assigned to a team.
func GetTeamHasUser(c context.Context, params *model.TeamUserParams) bool {
	return FromContext(c).GetTeamHasUser(params)
}

// CreateTeamUser assigns a user to a specific team.
func CreateTeamUser(c context.Context, params *model.TeamUserParams) error {
	return FromContext(c).CreateTeamUser(params)
}

// DeleteTeamUser removes a user from a specific team.
func DeleteTeamUser(c context.Context, params *model.TeamUserParams) error {
	return FromContext(c).DeleteTeamUser(params)
}

// GetTeamNamespaces retrieves namespaces for a team.
func GetTeamNamespaces(c context.Context, params *model.TeamNamespaceParams) (*model.Namespaces, error) {
	return FromContext(c).GetTeamNamespaces(params)
}

// GetTeamHasNamespace checks if a specific namespace is assigned to a team.
func GetTeamHasNamespace(c context.Context, params *model.TeamNamespaceParams) bool {
	return FromContext(c).GetTeamHasNamespace(params)
}

// CreateTeamNamespace assigns a namespace to a specific team.
func CreateTeamNamespace(c context.Context, params *model.TeamNamespaceParams) error {
	return FromContext(c).CreateTeamNamespace(params)
}

// DeleteTeamNamespace removes a namespace from a specific team.
func DeleteTeamNamespace(c context.Context, params *model.TeamNamespaceParams) error {
	return FromContext(c).DeleteTeamNamespace(params)
}
