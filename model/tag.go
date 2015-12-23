package model

import (
	"github.com/russross/meddler"
)

type Tag struct {
	ID           int64  `json:"id" meddler:"id,pk"`
	RepositoryID int64  `json:"-" meddler:"repository_id"`
	UserID       int64  `json:"-" meddler:"user_id"`
	Name         string `json:"name" meddler:"name"`
	Created      int64  `json:"created_at" meddler:"created_at"`
	Updated      int64  `json:"updated_at" meddler:"updated_at"`

	Repository *Repository `json:"repository"`
	User       *User       `json:"user"`
}
