package model

import (
	"fmt"
	"net/url"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// Registries is simply a collection of registry structs.
type Registries []*Registry

// Registry represents a registry model definition.
type Registry struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"name" sql:"unique_index"`
	Host      string    `json:"host" sql:"unique_index"`
	PlainHost string    `json:"-" sql:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Orgs      Orgs      `json:"orgs,omitempty"`
}

// AfterFind invokes required after loading a record from the database.
func (u *Registry) AfterFind(db *gorm.DB) {
	parsedURL, _ := url.Parse(u.Host)
	u.PlainHost = parsedURL.Host
}

// AfterSave invokes required actions after persisting.
func (u *Registry) AfterSave(db *gorm.DB) (err error) {
	orgs := Orgs{}

	err = db.Model(
		u,
	).Related(
		&orgs,
	).Error

	if err != nil {
		return err
	}

	for _, org := range orgs {
		err = db.Save(
			org,
		).Error

		if err != nil {
			return err
		}
	}

	return nil
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

// BeforeDelete invokes required actions before deletion.
func (u *Registry) BeforeDelete(tx *gorm.DB) error {
	orgs := Orgs{}

	tx.Model(
		u,
	).Related(
		&orgs,
	)

	if len(orgs) > 0 {
		return fmt.Errorf("Can't delete, still assigned to orgs.")
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

	if !govalidator.IsRequestURL(u.Host) {
		db.AddError(fmt.Errorf(
			"Host must be a valid URL",
		))
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
