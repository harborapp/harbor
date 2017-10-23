package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/pkg/model"
	"golang.org/x/net/context"
)

// GetRepos retrieves all available repos from the database.
func GetRepos(c context.Context) (*model.Repos, error) {
	return FromContext(c).GetRepos()
}

// CreateRepo creates a new repo.
func CreateRepo(c context.Context, record *model.Repo) error {
	return FromContext(c).CreateRepo(record, Current(c))
}

// UpdateRepo updates a repo.
func UpdateRepo(c context.Context, record *model.Repo) error {
	return FromContext(c).UpdateRepo(record, Current(c))
}

// DeleteRepo deletes a repo.
func DeleteRepo(c context.Context, record *model.Repo) error {
	return FromContext(c).DeleteRepo(record, Current(c))
}

// GetRepo retrieves a specific repo from the database.
func GetRepo(c context.Context, id string) (*model.Repo, *gorm.DB) {
	return FromContext(c).GetRepo(id)
}
