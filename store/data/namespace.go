package data

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetNamespaces retrieves all available namespaces from the database.
func (db *data) GetNamespaces() (*model.Namespaces, error) {
	records := &model.Namespaces{}

	err := db.Order(
		"name ASC",
	).Preload(
		"Registry",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).Find(
		&records,
	).Error

	return records, err
}

// CreateNamespace creates a new namespace.
func (db *data) CreateNamespace(record *model.Namespace) error {
	return db.Create(
		&record,
	).Error
}

// UpdateNamespace updates a namespace.
func (db *data) UpdateNamespace(record *model.Namespace) error {
	return db.Save(
		&record,
	).Error
}

// DeleteNamespace deletes a namespace.
func (db *data) DeleteNamespace(record *model.Namespace) error {
	return db.Delete(
		&record,
	).Error
}

// GetNamespace retrieves a specific namespace from the database.
func (db *data) GetNamespace(id string) (*model.Namespace, *gorm.DB) {
	record := &model.Namespace{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		&record,
	).Preload(
		"Registry",
	).Preload(
		"Users",
	).Preload(
		"Teams",
	).First(
		&record,
	)

	return record, res
}

// GetNamespaceUsers retrieves users for a namespace.
func (db *data) GetNamespaceUsers(params *model.NamespaceUserParams) (*model.Users, error) {
	namespace, _ := db.GetNamespace(params.Namespace)

	records := &model.Users{}

	err := db.Model(
		namespace,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetNamespaceHasUser checks if a specific user is assigned to a namespace.
func (db *data) GetNamespaceHasUser(params *model.NamespaceUserParams) bool {
	namespace, _ := db.GetNamespace(params.Namespace)
	user, _ := db.GetUser(params.User)

	count := db.Model(
		namespace,
	).Association(
		"Users",
	).Find(
		user,
	).Count()

	return count > 0
}

func (db *data) CreateNamespaceUser(params *model.NamespaceUserParams) error {
	namespace, _ := db.GetNamespace(params.Namespace)
	user, _ := db.GetUser(params.User)

	return db.Model(
		namespace,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeleteNamespaceUser(params *model.NamespaceUserParams) error {
	namespace, _ := db.GetNamespace(params.Namespace)
	user, _ := db.GetUser(params.User)

	return db.Model(
		namespace,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetNamespaceTeams retrieves teams for a namespace.
func (db *data) GetNamespaceTeams(params *model.NamespaceTeamParams) (*model.Teams, error) {
	namespace, _ := db.GetNamespace(params.Namespace)

	records := &model.Teams{}

	err := db.Model(
		namespace,
	).Association(
		"Teams",
	).Find(
		records,
	).Error

	return records, err
}

// GetNamespaceHasTeam checks if a specific team is assigned to a namespace.
func (db *data) GetNamespaceHasTeam(params *model.NamespaceTeamParams) bool {
	namespace, _ := db.GetNamespace(params.Namespace)
	team, _ := db.GetTeam(params.Team)

	count := db.Model(
		namespace,
	).Association(
		"Teams",
	).Find(
		team,
	).Count()

	return count > 0
}

func (db *data) CreateNamespaceTeam(params *model.NamespaceTeamParams) error {
	namespace, _ := db.GetNamespace(params.Namespace)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		namespace,
	).Association(
		"Teams",
	).Append(
		team,
	).Error
}

func (db *data) DeleteNamespaceTeam(params *model.NamespaceTeamParams) error {
	namespace, _ := db.GetNamespace(params.Namespace)
	team, _ := db.GetTeam(params.Team)

	return db.Model(
		namespace,
	).Association(
		"Teams",
	).Delete(
		team,
	).Error
}
