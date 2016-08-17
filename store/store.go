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
	CreateRegistry(*model.Registry) error

	// UpdateRegistry updates a registry.
	UpdateRegistry(*model.Registry) error

	// DeleteRegistry deletes a registry.
	DeleteRegistry(*model.Registry) error

	// GetRegistry retrieves a specific registry from the database.
	GetRegistry(string) (*model.Registry, *gorm.DB)

	// GetTags retrieves all available tags from the database.
	GetTags() (*model.Tags, error)

	// CreateTag creates a new tag.
	CreateTag(*model.Tag) error

	// UpdateTag updates a tag.
	UpdateTag(*model.Tag) error

	// DeleteTag deletes a tag.
	DeleteTag(*model.Tag) error

	// GetTag retrieves a specific tag from the database.
	GetTag(string) (*model.Tag, *gorm.DB)

	// GetRepositories retrieves all available repositories from the database.
	GetRepositories() (*model.Repositories, error)

	// CreateRepository creates a new repository.
	CreateRepository(*model.Repository) error

	// UpdateRepository updates a repository.
	UpdateRepository(*model.Repository) error

	// DeleteRepository deletes a repository.
	DeleteRepository(*model.Repository) error

	// GetRepository retrieves a specific repository from the database.
	GetRepository(string) (*model.Repository, *gorm.DB)

	// GetOrgs retrieves all available orgs from the database.
	GetOrgs() (*model.Orgs, error)

	// CreateOrg creates a new org.
	CreateOrg(*model.Org) error

	// UpdateOrg updates a org.
	UpdateOrg(*model.Org) error

	// DeleteOrg deletes a org.
	DeleteOrg(*model.Org) error

	// GetOrg retrieves a specific org from the database.
	GetOrg(string) (*model.Org, *gorm.DB)

	// GetOrgTeams retrieves teams for a org.
	GetOrgTeams(*model.OrgTeamParams) (*model.Teams, error)

	// GetOrgHasTeam checks if a specific team is assigned to a org.
	GetOrgHasTeam(*model.OrgTeamParams) bool

	// CreateOrgTeam assigns a team to a specific org.
	CreateOrgTeam(*model.OrgTeamParams) error

	// DeleteOrgTeam removes a team from a specific org.
	DeleteOrgTeam(*model.OrgTeamParams) error

	// GetOrgUsers retrieves users for a org.
	GetOrgUsers(*model.OrgUserParams) (*model.Users, error)

	// GetOrgHasUser checks if a specific user is assigned to a org.
	GetOrgHasUser(*model.OrgUserParams) bool

	// CreateOrgUser assigns a user to a specific org.
	CreateOrgUser(*model.OrgUserParams) error

	// DeleteOrgUser removes a user from a specific org.
	DeleteOrgUser(*model.OrgUserParams) error

	// GetUsers retrieves all available users from the database.
	GetUsers() (*model.Users, error)

	// CreateUser creates a new user.
	CreateUser(*model.User) error

	// UpdateUser updates a user.
	UpdateUser(*model.User) error

	// DeleteUser deletes a user.
	DeleteUser(*model.User) error

	// GetUser retrieves a specific user from the database.
	GetUser(string) (*model.User, *gorm.DB)

	// GetUserTeams retrieves teams for a user.
	GetUserTeams(*model.UserTeamParams) (*model.Teams, error)

	// GetUserHasTeam checks if a specific team is assigned to a user.
	GetUserHasTeam(*model.UserTeamParams) bool

	// CreateUserTeam assigns a team to a specific user.
	CreateUserTeam(*model.UserTeamParams) error

	// DeleteUserTeam removes a team from a specific user.
	DeleteUserTeam(*model.UserTeamParams) error

	// GetUserOrgs retrieves orgs for a user.
	GetUserOrgs(*model.UserOrgParams) (*model.Orgs, error)

	// GetUserHasOrg checks if a specific org is assigned to a user.
	GetUserHasOrg(*model.UserOrgParams) bool

	// CreateUserOrg assigns a org to a specific user.
	CreateUserOrg(*model.UserOrgParams) error

	// DeleteUserOrg removes a org from a specific user.
	DeleteUserOrg(*model.UserOrgParams) error

	// GetTeams retrieves all available teams from the database.
	GetTeams() (*model.Teams, error)

	// CreateTeam creates a new team.
	CreateTeam(*model.Team) error

	// UpdateTeam updates a team.
	UpdateTeam(*model.Team) error

	// DeleteTeam deletes a team.
	DeleteTeam(*model.Team) error

	// GetTeam retrieves a specific team from the database.
	GetTeam(string) (*model.Team, *gorm.DB)

	// GetTeamUsers retrieves users for a team.
	GetTeamUsers(*model.TeamUserParams) (*model.Users, error)

	// GetTeamHasUser checks if a specific user is assigned to a team.
	GetTeamHasUser(*model.TeamUserParams) bool

	// CreateTeamUser assigns a user to a specific team.
	CreateTeamUser(*model.TeamUserParams) error

	// DeleteTeamUser removes a user from a specific team.
	DeleteTeamUser(*model.TeamUserParams) error

	// GetTeamOrgs retrieves orgs for a team.
	GetTeamOrgs(*model.TeamOrgParams) (*model.Orgs, error)

	// GetTeamHasOrg checks if a specific org is assigned to a team.
	GetTeamHasOrg(*model.TeamOrgParams) bool

	// CreateTeamOrg assigns a org to a specific team.
	CreateTeamOrg(*model.TeamOrgParams) error

	// DeleteTeamOrg removes a org from a specific team.
	DeleteTeamOrg(*model.TeamOrgParams) error
}
