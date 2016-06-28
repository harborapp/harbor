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
