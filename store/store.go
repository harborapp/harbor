package store

import (
	"github.com/harborapp/harbor-api/model"
	"github.com/jinzhu/gorm"
)

//go:generate mockery -all -case=underscore

// Store implements all required data-layer functions for Harbor.
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
}
