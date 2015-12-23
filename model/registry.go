package model

import (
	"github.com/russross/meddler"
)

type Registry struct {
	ID      int64  `json:"id" meddler:"id,pk"`
	Name    string `json:"name" meddler:"name"`
	Host    string `json:"host" meddler:"host"`
	UseSSL  bool   `json:"use_ssl" meddler:"use_ssl"`
	Created int64  `json:"created_at" meddler:"created_at"`
	Updated int64  `json:"updated_at" meddler:"updated_at"`

	Namespaces []*Namespace `json:"namespaces"`
}
