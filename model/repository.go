package model

import (
	"github.com/russross/meddler"
)

type Repository struct {
	ID          int64  `json:"id" meddler:"id,pk"`
	NamespaceID int64  `json:"-" meddler:"namespace_id"`
	Name        string `json:"name" meddler:"name"`
	Created     int64  `json:"created_at" meddler:"created_at"`
	Updated     int64  `json:"updated_at" meddler:"updated_at"`

	Namespace *Namespace `json:"namespace"`
	Tags      []*Tag     `json:"tags"`
}
