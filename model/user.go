package model

import (
	"github.com/russross/meddler"
)

type User struct {
	ID       int64  `json:"id" meddler:"id,pk"`
	Email    string `json:"email" meddler:"email"`
	Username string `json:"username" meddler:"username"`
	Password string `json:"password" meddler:"password"`
	Admin    bool   `json:"admin" meddler:"admin"`
	Active   bool   `json:"active" meddler:"active"`
	Created  int64  `json:"created_at" meddler:"created_at"`
	Updated  int64  `json:"updated_at" meddler:"updated_at"`

	Teams []*Team `json:"teams"`
	Tags  []*Tag  `json:"tags"`
}
