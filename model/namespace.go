package model

import (
	"github.com/russross/meddler"
)

type User struct {
	ID         int64  `json:"id" meddler:"id,pk"`
	TeamID     int64  `json:"-" meddler:"team_id"`
	RegistryID int64  `json:"-" meddler:"registry_id"`
	Name       string `json:"name" meddler:"name"`
	Public     bool   `json:"public" meddler:"public"`
	Created    int64  `json:"created_at" meddler:"created_at"`
	Updated    int64  `json:"updated_at" meddler:"updated_at"`

	Team         *Team         `json:"team"`
	Registry     *Registry     `json:"registry"`
	Repositories []*Repository `json:"repositories"`
}
