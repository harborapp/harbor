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

// GetTeamOrgs retrieves orgs for a team.
func GetTeamOrgs(c context.Context, params *model.TeamOrgParams) (*model.Orgs, error) {
	return FromContext(c).GetTeamOrgs(params)
}

// GetTeamHasOrg checks if a specific org is assigned to a team.
func GetTeamHasOrg(c context.Context, params *model.TeamOrgParams) bool {
	return FromContext(c).GetTeamHasOrg(params)
}

// CreateTeamOrg assigns a org to a specific team.
func CreateTeamOrg(c context.Context, params *model.TeamOrgParams) error {
	return FromContext(c).CreateTeamOrg(params)
}

// DeleteTeamOrg removes a org from a specific team.
func DeleteTeamOrg(c context.Context, params *model.TeamOrgParams) error {
	return FromContext(c).DeleteTeamOrg(params)
}
