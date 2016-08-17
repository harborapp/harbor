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

// NamespaceUserParams represents the parameters to connect users with namespaces.
type NamespaceUserParams struct {
	Namespace string `json:"namespace"`
	User      string `json:"user"`
}

// UserNamespaceParams represents the parameters to connect namespaces with users.
type UserNamespaceParams struct {
	User      string `json:"user"`
	Namespace string `json:"namespace"`
}

// NamespaceTeamParams represents the parameters to connect teams with namespaces.
type NamespaceTeamParams struct {
	Namespace string `json:"namespace"`
	Team      string `json:"team"`
}

// TeamNamespaceParams represents the parameters to connect namespaces with teams.
type TeamNamespaceParams struct {
	Team      string `json:"team"`
	Namespace string `json:"namespace"`
}
