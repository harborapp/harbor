package storage

import (
	"errors"

	"github.com/umschlag/umschlag-api/pkg/model"
)

var (
	// ErrInvalidUserOrgPerm defines the error for invalid permissions.
	ErrInvalidUserOrgPerm = errors.New("invalid permission, can be user, admin or owner")
)

// GetUserOrgs retrieves orgs for an user.
func (db *data) GetUserOrgs(params *model.UserOrgParams) (*model.UserOrgs, error) {
	user, _ := db.GetUser(params.User)
	records := &model.UserOrgs{}

	err := db.engine.Where(
		"user_id = ?",
		user.ID,
	).Model(
		&model.UserOrg{},
	).Preload(
		"User",
	).Preload(
		"Org",
	).Find(
		records,
	).Error

	return records, err
}

// GetUserHasOrg checks if an specific org is assigned to an user.
func (db *data) GetUserHasOrg(params *model.UserOrgParams) bool {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	res := db.engine.Model(
		user,
	).Association(
		"Orgs",
	).Find(
		org,
	).Error

	return res == nil
}

// CreateUserOrg assigns an org to an specific user.
func (db *data) CreateUserOrg(params *model.UserOrgParams) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.engine.Create(
				&model.UserOrg{
					UserID: user.ID,
					OrgID:  org.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return ErrInvalidUserOrgPerm
}

// UpdateUserOrg updates the user org permission.
func (db *data) UpdateUserOrg(params *model.UserOrgParams) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	return db.engine.Model(
		&model.UserOrg{},
	).Where(
		"user_id = ? AND org_id = ?",
		user.ID,
		org.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

// DeleteUserOrg removes an org from an specific user.
func (db *data) DeleteUserOrg(params *model.UserOrgParams) error {
	user, _ := db.GetUser(params.User)
	org, _ := db.GetOrg(params.Org)

	return db.engine.Model(
		user,
	).Association(
		"Orgs",
	).Delete(
		org,
	).Error
}
