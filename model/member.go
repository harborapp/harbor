package model

import (
	"github.com/russross/meddler"
)

type Member struct {
	TeamID  int64 `json:"-" meddler:"team_id,pk"`
	UserID  int64 `json:"-" meddler:"user_id,pk"`
	Created int64 `json:"created_at" meddler:"created_at"`

	Team *Team `json:"team"`
	User *User `json:"user"`
}
