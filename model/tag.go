package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Tags is simply a collection of tag structs.
type Tags []*Tag

// Tag represents a tag model definition.
type Tag struct {
	ID           int         `json:"id" gorm:"primary_key"`
	Repository   *Repository `json:"repository,omitempty"`
	RepositoryID int         `json:"repository_id" sql:"index"`
	Slug         string      `json:"slug"`
	Name         string      `json:"name"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// BeforeSave invokes required actions before persisting.
func (u *Tag) BeforeSave(db *gorm.DB) (err error) {
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
				&Tag{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// AfterDelete invokes required actions after deletion.
func (u *Tag) AfterDelete(tx *gorm.DB) error {
	return nil
}

// Validate does some validation to be able to store the record.
func (u *Tag) Validate(db *gorm.DB) {
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
			&Tag{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
