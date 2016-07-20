package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

//go:generate mockery -all -case=underscore

// Store implements all required data-layer functions for Umschlag.
type Store interface {
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
}
