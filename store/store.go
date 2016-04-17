package store

import (
	"github.com/harborapp/harbor-api/model"
	"github.com/jinzhu/gorm"
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

// GetUsers retrieves all available users from the database.
func GetUsers(c context.Context) (*model.Users, error) {
	return FromContext(c).GetUsers()
}

// CreateUser creates a new user.
func CreateUser(c context.Context, record *model.User) error {
	return FromContext(c).CreateUser(record)
}

// UpdateUser updates a user.
func UpdateUser(c context.Context, record *model.User) error {
	return FromContext(c).UpdateUser(record)
}

// DeleteUser deletes a user.
func DeleteUser(c context.Context, record *model.User) error {
	return FromContext(c).DeleteUser(record)
}

// GetUser retrieves a specific user from the database.
func GetUser(c context.Context, id string) (*model.User, *gorm.DB) {
	return FromContext(c).GetUser(id)
}
