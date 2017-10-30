package storage

import (
	"errors"

	"github.com/umschlag/umschlag-api/pkg/model"
)

var (
	// ErrInvalidOrgUserPerm defines the error for invalid permissions.
	ErrInvalidOrgUserPerm = errors.New("invalid permission, can be user, admin or owner")
)

// GetOrgUsers retrieves users for an org.
func (db *data) GetOrgUsers(params *model.OrgUserParams) (*model.UserOrgs, error) {
	org, _ := db.GetOrg(params.Org)
	records := &model.UserOrgs{}

	err := db.engine.Where(
		"org_id = ?",
		org.ID,
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

// GetOrgHasUser checks if an specific user is assigned to an org.
func (db *data) GetOrgHasUser(params *model.OrgUserParams) bool {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	res := db.engine.Model(
		org,
	).Association(
		"Users",
	).Find(
		user,
	).Error

	return res == nil
}

// CreateOrgUser assigns an user to an specific org.
func (db *data) CreateOrgUser(params *model.OrgUserParams) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	for _, perm := range []string{"user", "admin", "owner"} {
		if params.Perm == perm {
			return db.engine.Create(
				&model.UserOrg{
					OrgID:  org.ID,
					UserID: user.ID,
					Perm:   params.Perm,
				},
			).Error
		}
	}

	return ErrInvalidOrgUserPerm
}

// UpdateOrgUser updates the org user permission.
func (db *data) UpdateOrgUser(params *model.OrgUserParams) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	return db.engine.Model(
		&model.UserOrg{},
	).Where(
		"org_id = ? AND user_id = ?",
		org.ID,
		user.ID,
	).Update(
		"perm",
		params.Perm,
	).Error
}

// DeleteOrgUser removes an user from an specific org.
func (db *data) DeleteOrgUser(params *model.OrgUserParams) error {
	org, _ := db.GetOrg(params.Org)
	user, _ := db.GetUser(params.User)

	return db.engine.Model(
		org,
	).Association(
		"Users",
	).Delete(
		user,
	).Error
}
