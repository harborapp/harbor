package model

import (
	"github.com/russross/meddler"
)

type Team struct {
	ID      int64  `json:"id" meddler:"id,pk"`
	Name    string `json:"name" meddler:"name"`
	Hidden  bool   `json:"hidden" meddler:"hidden"`
	Created int64  `json:"created_at" meddler:"created_at"`
	Updated int64  `json:"updated_at" meddler:"updated_at"`

	Users      []*User      `json:"users"`
	Namespaces []*Namespace `json:"namespaces"`
}
