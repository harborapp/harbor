package model

// TeamUserParams represents the parameters to connect users with teams.
type TeamUserParams struct {
	Team string `json:"team"`
	User string `json:"user"`
}

// UserTeamParams represents the parameters to connect teams with users.
type UserTeamParams struct {
	User string `json:"user"`
	Team string `json:"team"`
}

// OrgUserParams represents the parameters to connect users with orgs.
type OrgUserParams struct {
	Org  string `json:"org"`
	User string `json:"user"`
}

// UserOrgParams represents the parameters to connect orgs with users.
type UserOrgParams struct {
	User string `json:"user"`
	Org  string `json:"org"`
}

// OrgTeamParams represents the parameters to connect teams with orgs.
type OrgTeamParams struct {
	Org  string `json:"org"`
	Team string `json:"team"`
}

// TeamOrgParams represents the parameters to connect orgs with teams.
type TeamOrgParams struct {
	Team string `json:"team"`
	Org  string `json:"org"`
}
