package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
	"golang.org/x/net/context"
)

// GetNamespaces retrieves all available namespaces from the database.
func GetNamespaces(c context.Context) (*model.Namespaces, error) {
	return FromContext(c).GetNamespaces()
}

// CreateNamespace creates a new namespace.
func CreateNamespace(c context.Context, record *model.Namespace) error {
	return FromContext(c).CreateNamespace(record)
}

// UpdateNamespace updates a namespace.
func UpdateNamespace(c context.Context, record *model.Namespace) error {
	return FromContext(c).UpdateNamespace(record)
}

// DeleteNamespace deletes a namespace.
func DeleteNamespace(c context.Context, record *model.Namespace) error {
	return FromContext(c).DeleteNamespace(record)
}

// GetNamespace retrieves a specific namespace from the database.
func GetNamespace(c context.Context, id string) (*model.Namespace, *gorm.DB) {
	return FromContext(c).GetNamespace(id)
}

// GetNamespaceUsers retrieves users for a namespace.
func GetNamespaceUsers(c context.Context, params *model.NamespaceUserParams) (*model.Users, error) {
	return FromContext(c).GetNamespaceUsers(params)
}

// GetNamespaceHasUser checks if a specific user is assigned to a namespace.
func GetNamespaceHasUser(c context.Context, params *model.NamespaceUserParams) bool {
	return FromContext(c).GetNamespaceHasUser(params)
}

// CreateNamespaceUser assigns a user to a specific namespace.
func CreateNamespaceUser(c context.Context, params *model.NamespaceUserParams) error {
	return FromContext(c).CreateNamespaceUser(params)
}

// DeleteNamespaceUser removes a user from a specific namespace.
func DeleteNamespaceUser(c context.Context, params *model.NamespaceUserParams) error {
	return FromContext(c).DeleteNamespaceUser(params)
}

// GetNamespaceTeams retrieves teams for a namespace.
func GetNamespaceTeams(c context.Context, params *model.NamespaceTeamParams) (*model.Teams, error) {
	return FromContext(c).GetNamespaceTeams(params)
}

// GetNamespaceHasTeam checks if a specific team is assigned to a namespace.
func GetNamespaceHasTeam(c context.Context, params *model.NamespaceTeamParams) bool {
	return FromContext(c).GetNamespaceHasTeam(params)
}

// CreateNamespaceTeam assigns a team to a specific namespace.
func CreateNamespaceTeam(c context.Context, params *model.NamespaceTeamParams) error {
	return FromContext(c).CreateNamespaceTeam(params)
}

// DeleteNamespaceTeam removes a team from a specific namespace.
func DeleteNamespaceTeam(c context.Context, params *model.NamespaceTeamParams) error {
	return FromContext(c).DeleteNamespaceTeam(params)
}
