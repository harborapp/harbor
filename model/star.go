package model

import (
	"github.com/russross/meddler"
)

type Star struct {
	UserID       int64 `json:"-" meddler:"user_id,pk"`
	RepositoryID int64 `json:"-" meddler:"repository_id,pk"`
	Created      int64 `json:"created_at" meddler:"created_at"`

	User       *User       `json:"user"`
	Repository *Repository `json:"repository"`
}
