package storage

import (
	"github.com/umschlag/umschlag-api/pkg/model"
)

// Store implements all required database-layer functions.
type Store interface {
	// GetRegistries retrieves all available registries from storage.
	GetRegistries() (*model.Registries, error)

	// CreateRegistry creates an new registry.
	CreateRegistry(*model.Registry) error

	// UpdateRegistry updates an registry.
	UpdateRegistry(*model.Registry) error

	// DeleteRegistry deletes an registry.
	DeleteRegistry(*model.Registry) error

	// GetRegistry retrieves a specific registry from storage.
	GetRegistry(string) (*model.Registry, error)

	// GetRepos retrieves all available repos from storage.
	GetRepos(*model.ReposFilter) (*model.Repos, error)

	// CreateRepo creates an new repo.
	CreateRepo(*model.Repo) error

	// UpdateRepo updates an repo.
	UpdateRepo(*model.Repo) error

	// DeleteRepo deletes an repo.
	DeleteRepo(*model.Repo) error

	// GetRepo retrieves a specific repo from storage.
	GetRepo(string) (*model.Repo, error)

	// GetTags retrieves all available tags from storage.
	GetTags(*model.TagsFilter) (*model.Tags, error)

	// CreateTag creates an new tag.
	CreateTag(*model.Tag) error

	// UpdateTag updates an tag.
	UpdateTag(*model.Tag) error

	// DeleteTag deletes an tag.
	DeleteTag(*model.Tag) error

	// GetTag retrieves a specific tag from storage.
	GetTag(string) (*model.Tag, error)

	// GetOrgs retrieves all available orgs from storage.
	GetOrgs() (*model.Orgs, error)

	// CreateOrg creates an new org.
	CreateOrg(*model.Org) error

	// UpdateOrg updates an org.
	UpdateOrg(*model.Org) error

	// DeleteOrg deletes an org.
	DeleteOrg(*model.Org) error

	// GetOrg retrieves a specific org from storage.
	GetOrg(string) (*model.Org, error)

	// GetOrgTeams retrieves teams for an org.
	GetOrgTeams(*model.OrgTeamParams) (*model.TeamOrgs, error)

	// GetOrgHasTeam checks if an specific team is assigned to an org.
	GetOrgHasTeam(*model.OrgTeamParams) bool

	// CreateOrgTeam assigns a team to an specific org.
	CreateOrgTeam(*model.OrgTeamParams) error

	// UpdateOrgTeam updates the org team permission.
	UpdateOrgTeam(*model.OrgTeamParams) error

	// DeleteOrgTeam removes a team from an specific org.
	DeleteOrgTeam(*model.OrgTeamParams) error

	// GetOrgUsers retrieves users for an org.
	GetOrgUsers(*model.OrgUserParams) (*model.UserOrgs, error)

	// GetOrgHasUser checks if an specific user is assigned to an org.
	GetOrgHasUser(*model.OrgUserParams) bool

	// CreateOrgUser assigns an user to an specific org.
	CreateOrgUser(*model.OrgUserParams) error

	// UpdateOrgUser updates the org user permission.
	UpdateOrgUser(*model.OrgUserParams) error

	// DeleteOrgUser removes an user from an specific org.
	DeleteOrgUser(*model.OrgUserParams) error

	// GetTeams retrieves all available teams from storage.
	GetTeams() (*model.Teams, error)

	// CreateTeam creates an new team.
	CreateTeam(*model.Team) error

	// UpdateTeam updates an team.
	UpdateTeam(*model.Team) error

	// DeleteTeam deletes an team.
	DeleteTeam(*model.Team) error

	// GetTeam retrieves a specific team from storage.
	GetTeam(string) (*model.Team, error)

	// GetTeamOrgs retrieves orgs for a team.
	GetTeamOrgs(*model.TeamOrgParams) (*model.TeamOrgs, error)

	// GetTeamHasOrg checks if an specific org is assigned to a team.
	GetTeamHasOrg(*model.TeamOrgParams) bool

	// CreateTeamOrg assigns an org to a specific team.
	CreateTeamOrg(*model.TeamOrgParams) error

	// UpdateTeamOrg updates the team org permission.
	UpdateTeamOrg(*model.TeamOrgParams) error

	// DeleteTeamOrg removes an org from a specific team.
	DeleteTeamOrg(*model.TeamOrgParams) error

	// GetTeamUsers retrieves users for a team.
	GetTeamUsers(*model.TeamUserParams) (*model.TeamUsers, error)

	// GetTeamHasUser checks if an specific user is assigned to a team.
	GetTeamHasUser(*model.TeamUserParams) bool

	// CreateTeamUser assigns an user to a specific team.
	CreateTeamUser(*model.TeamUserParams) error

	// UpdateTeamUser updates the team user permission.
	UpdateTeamUser(*model.TeamUserParams) error

	// DeleteTeamUser removes an user from a specific team.
	DeleteTeamUser(*model.TeamUserParams) error

	// GetUsers retrieves all available users from storage.
	GetUsers() (*model.Users, error)

	// CreateUser creates an new user.
	CreateUser(*model.User) error

	// UpdateUser updates an user.
	UpdateUser(*model.User) error

	// DeleteUser deletes an user.
	DeleteUser(*model.User) error

	// GetUser retrieves a specific user from storage.
	GetUser(string) (*model.User, error)

	// GetUserOrgs retrieves orgs for an user.
	GetUserOrgs(*model.UserOrgParams) (*model.UserOrgs, error)

	// GetUserHasOrg checks if an specific org is assigned to an user.
	GetUserHasOrg(*model.UserOrgParams) bool

	// CreateUserOrg assigns an org to an specific user.
	CreateUserOrg(*model.UserOrgParams) error

	// UpdateUserOrg updates the user org permission.
	UpdateUserOrg(*model.UserOrgParams) error

	// DeleteUserOrg removes an org from an specific user.
	DeleteUserOrg(*model.UserOrgParams) error

	// GetUserTeams retrieves teams for an user.
	GetUserTeams(*model.UserTeamParams) (*model.TeamUsers, error)

	// GetUserHasTeam checks if a specific team is assigned to an user.
	GetUserHasTeam(*model.UserTeamParams) bool

	// CreateUserTeam assigns a team to an specific user.
	CreateUserTeam(*model.UserTeamParams) error

	// UpdateUserTeam updates the user team permission.
	UpdateUserTeam(*model.UserTeamParams) error

	// DeleteUserTeam removes a team from an specific user.
	DeleteUserTeam(*model.UserTeamParams) error
}
