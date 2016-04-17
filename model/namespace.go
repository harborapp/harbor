package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Namespaces is simply a collection of namespace structs.
type Namespaces []*Namespace

// Namespace represents a namespace model definition.
type Namespace struct {
	ID           int          `json:"id" gorm:"primary_key"`
	Registry     *Registry    `json:"registry,omitempty"`
	RegistryID   int          `json:"registry_id" sql:"index"`
	Slug         string       `json:"slug"`
	Name         string       `json:"name"`
	Public       bool         `json:"private" sql:"default:false"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Repositories Repositories `json:"repositories,omitempty"`
	Teams        Teams        `json:"teams,omitempty" gorm:"many2many:team_namespaces;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Namespace) BeforeSave(db *gorm.DB) (err error) {
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
				&Namespace{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Namespace) Validate(db *gorm.DB) {
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
			&Namespace{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
