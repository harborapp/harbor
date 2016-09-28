package model

import (
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Repos is simply a collection of repo structs.
type Repos []*Repo

// Repo represents a repo model definition.
type Repo struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Org       *Org      `json:"org,omitempty"`
	OrgID     int64     `json:"org_id" sql:"index"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	Public    bool      `json:"public" sql:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tags      Tags      `json:"tags,omitempty"`
}

// UpdateFullName updates the cached full name of the repo
func (u *Repo) UpdateFullName(db *gorm.DB) (err error) {
	var (
		org      = &Org{}
		registry = &Registry{}
		handle   *gorm.DB
	)

	if handle = db.First(org, u.OrgID); handle.RecordNotFound() {
		return fmt.Errorf("Failed to find parent org")
	}

	if handle = db.First(registry, org.RegistryID); handle.RecordNotFound() {
		return fmt.Errorf("Failed to find parent registry")
	}

	u.FullName = fmt.Sprintf(
		"%s/%s/%s",
		registry.PlainHost,
		org.Name,
		u.Name,
	)

	return nil
}

// AfterSave invokes required actions after persisting.
func (u *Repo) AfterSave(db *gorm.DB) (err error) {
	tags := Tags{}

	err = db.Model(
		u,
	).Related(
		&tags,
	).Error

	if err != nil {
		return err
	}

	for _, tag := range tags {
		err = db.Save(
			tag,
		).Error

		if err != nil {
			return err
		}
	}

	return nil
}

// BeforeSave invokes required actions before persisting.
func (u *Repo) BeforeSave(db *gorm.DB) (err error) {
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
				&Repo{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return u.UpdateFullName(db)
}

// BeforeDelete invokes required actions before deletion.
func (u *Repo) BeforeDelete(tx *gorm.DB) error {
	tags := Tags{}

	tx.Model(
		u,
	).Related(
		&tags,
	)

	for _, tag := range tags {
		if err := tx.Delete(tag).Error; err != nil {
			return err
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Repo) Validate(db *gorm.DB) {
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
			&Repo{},
		).RecordNotFound()

		if !notFound {
			db.AddError(fmt.Errorf("Name is already present"))
		}
	}
}
