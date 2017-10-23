package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
	"golang.org/x/net/context"
)

// GetOrgs retrieves all available orgs from the database.
func GetOrgs(c context.Context) (*model.Orgs, error) {
	return FromContext(c).GetOrgs()
}

// CreateOrg creates a new org.
func CreateOrg(c context.Context, record *model.Org) error {
	return FromContext(c).CreateOrg(record, Current(c))
}

// UpdateOrg updates a org.
func UpdateOrg(c context.Context, record *model.Org) error {
	return FromContext(c).UpdateOrg(record, Current(c))
}

// DeleteOrg deletes a org.
func DeleteOrg(c context.Context, record *model.Org) error {
	return FromContext(c).DeleteOrg(record, Current(c))
}

// GetOrg retrieves a specific org from the database.
func GetOrg(c context.Context, id string) (*model.Org, *gorm.DB) {
	return FromContext(c).GetOrg(id)
}

// GetOrgUsers retrieves users for a org.
func GetOrgUsers(c context.Context, params *model.OrgUserParams) (*model.UserOrgs, error) {
	return FromContext(c).GetOrgUsers(params)
}

// GetOrgHasUser checks if a specific user is assigned to a org.
func GetOrgHasUser(c context.Context, params *model.OrgUserParams) bool {
	return FromContext(c).GetOrgHasUser(params)
}

// CreateOrgUser assigns a user to a specific org.
func CreateOrgUser(c context.Context, params *model.OrgUserParams) error {
	return FromContext(c).CreateOrgUser(params, Current(c))
}

// UpdateOrgUser updates the org user permission.
func UpdateOrgUser(c context.Context, params *model.OrgUserParams) error {
	return FromContext(c).UpdateOrgUser(params, Current(c))
}

// DeleteOrgUser removes a user from a specific org.
func DeleteOrgUser(c context.Context, params *model.OrgUserParams) error {
	return FromContext(c).DeleteOrgUser(params, Current(c))
}

// GetOrgTeams retrieves teams for a org.
func GetOrgTeams(c context.Context, params *model.OrgTeamParams) (*model.TeamOrgs, error) {
	return FromContext(c).GetOrgTeams(params)
}

// GetOrgHasTeam checks if a specific team is assigned to a org.
func GetOrgHasTeam(c context.Context, params *model.OrgTeamParams) bool {
	return FromContext(c).GetOrgHasTeam(params)
}

// CreateOrgTeam assigns a team to a specific org.
func CreateOrgTeam(c context.Context, params *model.OrgTeamParams) error {
	return FromContext(c).CreateOrgTeam(params, Current(c))
}

// UpdateOrgTeam updates the org team permission.
func UpdateOrgTeam(c context.Context, params *model.OrgTeamParams) error {
	return FromContext(c).UpdateOrgTeam(params, Current(c))
}

// DeleteOrgTeam removes a team from a specific org.
func DeleteOrgTeam(c context.Context, params *model.OrgTeamParams) error {
	return FromContext(c).DeleteOrgTeam(params, Current(c))
}
