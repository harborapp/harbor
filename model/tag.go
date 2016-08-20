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
	ID        int       `json:"id" gorm:"primary_key"`
	Repo      *Repo     `json:"repo,omitempty"`
	RepoID    int       `json:"repo_id" sql:"index"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	Public    bool      `json:"public" sql:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateFullName updates the cached full name of the tag
func (u *Tag) UpdateFullName(db *gorm.DB) (err error) {
	var (
		repo     = &Repo{}
		org      = &Org{}
		registry = &Registry{}
		handle   *gorm.DB
	)

	if handle = db.First(repo, u.RepoID); handle.RecordNotFound() {
		return fmt.Errorf("Failed to find parent repo")
	}

	if handle = db.First(org, repo.OrgID); handle.RecordNotFound() {
		return fmt.Errorf("Failed to find parent org")
	}

	if handle = db.First(registry, org.RegistryID); handle.RecordNotFound() {
		return fmt.Errorf("Failed to find parent registry")
	}

	u.FullName = fmt.Sprintf(
		"%s/%s/%s:%s",
		registry.Host,
		org.Name,
		repo.Name,
		u.Name,
	)

	return nil
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

	return u.UpdateFullName(db)
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
