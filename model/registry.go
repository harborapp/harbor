package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Registries is simply a collection of registry structs.
type Registries []*Registry

// Registry represents a registry model definition.
type Registry struct {
	ID         int        `json:"id" gorm:"primary_key"`
	Slug       string     `json:"slug" sql:"unique_index"`
	Name       string     `json:"name" sql:"unique_index"`
	Host       string     `json:"host" sql:"unique_index"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Namespaces Namespaces `json:"namespaces,omitempty"`
}

// BeforeSave invokes required actions before persisting.
func (u *Registry) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				u.Slug = slugify.Slugify(u.Name)
			} else {
				u.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", u.Name, i),
				)
			}

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Registry{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Registry) Validate(db *gorm.DB) {
	if !govalidator.StringLength(u.Name, "1", "255") {
		db.AddError(fmt.Errorf("Name should be longer than 1 and shorter than 255"))
	}

	if u.Name != "" {
		notFound := db.Where(
			"name = ?",
			u.Name,
		).Not(
			"id",
			u.ID,
		).First(
			&Registry{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}

	if !govalidator.StringLength(u.Host, "1", "255") {
		db.AddError(fmt.Errorf("Host should be longer than 1 and shorter than 255"))
	}

	if u.Host != "" {
		notFound := db.Where(
			"host = ?",
			u.Host,
		).Not(
			"id",
			u.ID,
		).First(
			&Registry{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Host is already present"))
		}
	}
}
