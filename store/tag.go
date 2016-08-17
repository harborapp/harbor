package store

import (
	"github.com/jinzhu/gorm"
	"github.com/umschlag/umschlag-api/model"
	"golang.org/x/net/context"
)

// GetTags retrieves all available tags from the database.
func GetTags(c context.Context) (*model.Tags, error) {
	return FromContext(c).GetTags()
}

// CreateTag creates a new tag.
func CreateTag(c context.Context, record *model.Tag) error {
	return FromContext(c).CreateTag(record)
}

// UpdateTag updates a tag.
func UpdateTag(c context.Context, record *model.Tag) error {
	return FromContext(c).UpdateTag(record)
}

// DeleteTag deletes a tag.
func DeleteTag(c context.Context, record *model.Tag) error {
	return FromContext(c).DeleteTag(record)
}

// GetTag retrieves a specific tag from the database.
func GetTag(c context.Context, id string) (*model.Tag, *gorm.DB) {
	return FromContext(c).GetTag(id)
}
