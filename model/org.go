package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Orgs is simply a collection of org structs.
type Orgs []*Org

// Org represents a org model definition.
type Org struct {
	ID           int          `json:"id" gorm:"primary_key"`
	Registry     *Registry    `json:"registry,omitempty"`
	RegistryID   int          `json:"registry_id" sql:"index"`
	Slug         string       `json:"slug"`
	Name         string       `json:"name"`
	Public       bool         `json:"private" sql:"default:false"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Repositories Repositories `json:"repositories,omitempty"`
	Teams        Teams        `json:"teams,omitempty" gorm:"many2many:team_orgs;"`
	Users        Users        `json:"users,omitempty" gorm:"many2many:user_orgs;"`
}

// BeforeSave invokes required actions before persisting.
func (u *Org) BeforeSave(db *gorm.DB) (err error) {
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
				&Org{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *Org) AfterDelete(tx *gorm.DB) error {
	for _, repository := range u.Repositories {
		if err := tx.Delete(repository).Error; err != nil {
			return err
		}
	}

	if err := tx.Model(u).Association("Teams").Clear().Error; err != nil {
		return err
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Org) Validate(db *gorm.DB) {
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
			&Org{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
