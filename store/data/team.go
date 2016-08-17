package data

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
)

// GetTeams retrieves all available teams from the database.
func (db *data) GetTeams() (*model.Teams, error) {
	records := &model.Teams{}

	err := db.Order(
		"name ASC",
	).Find(
		records,
	).Error

	return records, err
}

// CreateTeam creates a new team.
func (db *data) CreateTeam(record *model.Team) error {
	return db.Create(
		record,
	).Error
}

// UpdateTeam updates a team.
func (db *data) UpdateTeam(record *model.Team) error {
	return db.Save(
		record,
	).Error
}

// DeleteTeam deletes a team.
func (db *data) DeleteTeam(record *model.Team) error {
	return db.Delete(
		record,
	).Error
}

// GetTeam retrieves a specific team from the database.
func (db *data) GetTeam(id string) (*model.Team, *gorm.DB) {
	record := &model.Team{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).Model(
		record,
	).First(
		record,
	)

	return record, res
}

// GetTeamUsers retrieves users for a team.
func (db *data) GetTeamUsers(params *model.TeamUserParams) (*model.Users, error) {
	team, _ := db.GetTeam(params.Team)

	records := &model.Users{}

	err := db.Model(
		team,
	).Association(
		"Users",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasUser checks if a specific user is assigned to a team.
func (db *data) GetTeamHasUser(params *model.TeamUserParams) bool {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	count := db.Model(
		team,
	).Association(
		"Users",
	).Find(
		user,
	).Count()

	return count > 0
}

func (db *data) CreateTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		team,
	).Association(
		"Users",
	).Append(
		user,
	).Error
}

func (db *data) DeleteTeamUser(params *model.TeamUserParams) error {
	team, _ := db.GetTeam(params.Team)
	user, _ := db.GetUser(params.User)

	return db.Model(
		team,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}

// GetTeamNamespaces retrieves namespaces for a team.
func (db *data) GetTeamNamespaces(params *model.TeamNamespaceParams) (*model.Namespaces, error) {
	team, _ := db.GetTeam(params.Team)

	records := &model.Namespaces{}

	err := db.Model(
		team,
	).Association(
		"Namespaces",
	).Find(
		records,
	).Error

	return records, err
}

// GetTeamHasNamespace checks if a specific namespace is assigned to a team.
func (db *data) GetTeamHasNamespace(params *model.TeamNamespaceParams) bool {
	team, _ := db.GetTeam(params.Team)
	namespace, _ := db.GetNamespace(params.Namespace)

	count := db.Model(
		team,
	).Association(
		"Namespaces",
	).Find(
		namespace,
	).Count()

	return count > 0
}

func (db *data) CreateTeamNamespace(params *model.TeamNamespaceParams) error {
	team, _ := db.GetTeam(params.Team)
	namespace, _ := db.GetNamespace(params.Namespace)

	return db.Model(
		team,
	).Association(
		"Namespaces",
	).Append(
		namespace,
	).Error
}

func (db *data) DeleteTeamNamespace(params *model.TeamNamespaceParams) error {
	team, _ := db.GetTeam(params.Team)
	namespace, _ := db.GetNamespace(params.Namespace)

	return db.Model(
		team,
	).Association(
		"Namespaces",
	).Delete(
		namespace,
	).Error
}
