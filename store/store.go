package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

//go:generate mockery -all -case=underscore

// Store implements all required data-layer functions for Umschlag.
type Store interface {
	// GetRegistries retrieves all available registries from the database.
	GetRegistries() (*model.Registries, error)

	// CreateRegistry creates a new registry.
	CreateRegistry(*model.Registry, *model.User) error

	// UpdateRegistry updates a registry.
	UpdateRegistry(*model.Registry, *model.User) error

	// DeleteRegistry deletes a registry.
	DeleteRegistry(*model.Registry, *model.User) error

	// GetRegistry retrieves a specific registry from the database.
	GetRegistry(string) (*model.Registry, *gorm.DB)

	// GetTags retrieves all available tags from the database.
	GetTags() (*model.Tags, error)

	// CreateTag creates a new tag.
	CreateTag(*model.Tag, *model.User) error

	// UpdateTag updates a tag.
	UpdateTag(*model.Tag, *model.User) error

	// DeleteTag deletes a tag.
	DeleteTag(*model.Tag, *model.User) error

	// GetTag retrieves a specific tag from the database.
	GetTag(string) (*model.Tag, *gorm.DB)

	// GetRepos retrieves all available repos from the database.
	GetRepos() (*model.Repos, error)

	// CreateRepo creates a new repo.
	CreateRepo(*model.Repo, *model.User) error

	// UpdateRepo updates a repo.
	UpdateRepo(*model.Repo, *model.User) error

	// DeleteRepo deletes a repo.
	DeleteRepo(*model.Repo, *model.User) error

	// GetRepo retrieves a specific repo from the database.
	GetRepo(string) (*model.Repo, *gorm.DB)

	// GetOrgs retrieves all available orgs from the database.
	GetOrgs() (*model.Orgs, error)

	// CreateOrg creates a new org.
	CreateOrg(*model.Org, *model.User) error

	// UpdateOrg updates a org.
	UpdateOrg(*model.Org, *model.User) error

	// DeleteOrg deletes a org.
	DeleteOrg(*model.Org, *model.User) error

	// GetOrg retrieves a specific org from the database.
	GetOrg(string) (*model.Org, *gorm.DB)

	// GetOrgTeams retrieves teams for a org.
	GetOrgTeams(*model.OrgTeamParams) (*model.TeamOrgs, error)

	// GetOrgHasTeam checks if a specific team is assigned to a org.
	GetOrgHasTeam(*model.OrgTeamParams) bool

	// CreateOrgTeam assigns a team to a specific org.
	CreateOrgTeam(*model.OrgTeamParams, *model.User) error

	// UpdateOrgTeam updates the org team permission.
	UpdateOrgTeam(*model.OrgTeamParams, *model.User) error

	// DeleteOrgTeam removes a team from a specific org.
	DeleteOrgTeam(*model.OrgTeamParams, *model.User) error

	// GetOrgUsers retrieves users for a org.
	GetOrgUsers(*model.OrgUserParams) (*model.UserOrgs, error)

	// GetOrgHasUser checks if a specific user is assigned to a org.
	GetOrgHasUser(*model.OrgUserParams) bool

	// CreateOrgUser assigns a user to a specific org.
	CreateOrgUser(*model.OrgUserParams, *model.User) error

	// UpdateOrgUser updates the org user permission.
	UpdateOrgUser(*model.OrgUserParams, *model.User) error

	// DeleteOrgUser removes a user from a specific org.
	DeleteOrgUser(*model.OrgUserParams, *model.User) error

	// GetUsers retrieves all available users from the database.
	GetUsers() (*model.Users, error)

	// CreateUser creates a new user.
	CreateUser(*model.User, *model.User) error

	// UpdateUser updates a user.
	UpdateUser(*model.User, *model.User) error

	// DeleteUser deletes a user.
	DeleteUser(*model.User, *model.User) error

	// GetUser retrieves a specific user from the database.
	GetUser(string) (*model.User, *gorm.DB)

	// GetUserTeams retrieves teams for a user.
	GetUserTeams(*model.UserTeamParams) (*model.TeamUsers, error)

	// GetUserHasTeam checks if a specific team is assigned to a user.
	GetUserHasTeam(*model.UserTeamParams) bool

	// CreateUserTeam assigns a team to a specific user.
	CreateUserTeam(*model.UserTeamParams, *model.User) error

	// UpdateUserTeam updates the user team permission.
	UpdateUserTeam(*model.UserTeamParams, *model.User) error

	// DeleteUserTeam removes a team from a specific user.
	DeleteUserTeam(*model.UserTeamParams, *model.User) error

	// GetUserOrgs retrieves orgs for a user.
	GetUserOrgs(*model.UserOrgParams) (*model.UserOrgs, error)

	// GetUserHasOrg checks if a specific org is assigned to a user.
	GetUserHasOrg(*model.UserOrgParams) bool

	// CreateUserOrg assigns a org to a specific user.
	CreateUserOrg(*model.UserOrgParams, *model.User) error

	// UpdateUserOrg updates the user org permission.
	UpdateUserOrg(*model.UserOrgParams, *model.User) error

	// DeleteUserOrg removes a org from a specific user.
	DeleteUserOrg(*model.UserOrgParams, *model.User) error

	// GetTeams retrieves all available teams from the database.
	GetTeams() (*model.Teams, error)

	// CreateTeam creates a new team.
	CreateTeam(*model.Team, *model.User) error

	// UpdateTeam updates a team.
	UpdateTeam(*model.Team, *model.User) error

	// DeleteTeam deletes a team.
	DeleteTeam(*model.Team, *model.User) error

	// GetTeam retrieves a specific team from the database.
	GetTeam(string) (*model.Team, *gorm.DB)

	// GetTeamUsers retrieves users for a team.
	GetTeamUsers(*model.TeamUserParams) (*model.TeamUsers, error)

	// GetTeamHasUser checks if a specific user is assigned to a team.
	GetTeamHasUser(*model.TeamUserParams) bool

	// CreateTeamUser assigns a user to a specific team.
	CreateTeamUser(*model.TeamUserParams, *model.User) error

	// UpdateTeamUser updates the team user permission.
	UpdateTeamUser(*model.TeamUserParams, *model.User) error

	// DeleteTeamUser removes a user from a specific team.
	DeleteTeamUser(*model.TeamUserParams, *model.User) error

	// GetTeamOrgs retrieves orgs for a team.
	GetTeamOrgs(*model.TeamOrgParams) (*model.TeamOrgs, error)

	// GetTeamHasOrg checks if a specific org is assigned to a team.
	GetTeamHasOrg(*model.TeamOrgParams) bool

	// CreateTeamOrg assigns a org to a specific team.
	CreateTeamOrg(*model.TeamOrgParams, *model.User) error

	// UpdateTeamOrg updates the team org permission.
	UpdateTeamOrg(*model.TeamOrgParams, *model.User) error

	// DeleteTeamOrg removes a org from a specific team.
	DeleteTeamOrg(*model.TeamOrgParams, *model.User) error
}
