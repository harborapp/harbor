package model

// UserOrgs is simply a collection of user org structs.
type UserOrgs []*UserOrg

// UserOrg represents a user org model definition.
type UserOrg struct {
	UserID int64  `json:"user_id" sql:"index"`
	User   *User  `json:"user,omitempty"`
	OrgID  int64  `json:"org_id" sql:"index"`
	Org    *Org   `json:"org,omitempty"`
	Perm   string `json:"perm,omitempty"`
}
