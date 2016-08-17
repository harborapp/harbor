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

	// GetNamespaces retrieves all available namespaces from the database.
	GetNamespaces() (*model.Namespaces, error)

	// CreateNamespace creates a new namespace.
	CreateNamespace(*model.Namespace) error

	// UpdateNamespace updates a namespace.
	UpdateNamespace(*model.Namespace) error

	// DeleteNamespace deletes a namespace.
	DeleteNamespace(*model.Namespace) error

	// GetNamespace retrieves a specific namespace from the database.
	GetNamespace(string) (*model.Namespace, *gorm.DB)

	// GetNamespaceTeams retrieves teams for a namespace.
	GetNamespaceTeams(*model.NamespaceTeamParams) (*model.Teams, error)

	// GetNamespaceHasTeam checks if a specific team is assigned to a namespace.
	GetNamespaceHasTeam(*model.NamespaceTeamParams) bool

	// CreateNamespaceTeam assigns a team to a specific namespace.
	CreateNamespaceTeam(*model.NamespaceTeamParams) error

	// DeleteNamespaceTeam removes a team from a specific namespace.
	DeleteNamespaceTeam(*model.NamespaceTeamParams) error

	// GetNamespaceUsers retrieves users for a namespace.
	GetNamespaceUsers(*model.NamespaceUserParams) (*model.Users, error)

	// GetNamespaceHasUser checks if a specific user is assigned to a namespace.
	GetNamespaceHasUser(*model.NamespaceUserParams) bool

	// CreateNamespaceUser assigns a user to a specific namespace.
	CreateNamespaceUser(*model.NamespaceUserParams) error

	// DeleteNamespaceUser removes a user from a specific namespace.
	DeleteNamespaceUser(*model.NamespaceUserParams) error

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

	// GetUserNamespaces retrieves namespaces for a user.
	GetUserNamespaces(*model.UserNamespaceParams) (*model.Namespaces, error)

	// GetUserHasNamespace checks if a specific namespace is assigned to a user.
	GetUserHasNamespace(*model.UserNamespaceParams) bool

	// CreateUserNamespace assigns a namespace to a specific user.
	CreateUserNamespace(*model.UserNamespaceParams) error

	// DeleteUserNamespace removes a namespace from a specific user.
	DeleteUserNamespace(*model.UserNamespaceParams) error

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

	// GetTeamNamespaces retrieves namespaces for a team.
	GetTeamNamespaces(*model.TeamNamespaceParams) (*model.Namespaces, error)

	// GetTeamHasNamespace checks if a specific namespace is assigned to a team.
	GetTeamHasNamespace(*model.TeamNamespaceParams) bool

	// CreateTeamNamespace assigns a namespace to a specific team.
	CreateTeamNamespace(*model.TeamNamespaceParams) error

	// DeleteTeamNamespace removes a namespace from a specific team.
	DeleteTeamNamespace(*model.TeamNamespaceParams) error
}
