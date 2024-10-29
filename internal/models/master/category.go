package master

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Category) TableName() string {
	return "master_categories"
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	return c.validate()
}

func (c *Category) BeforeUpdate(tx *gorm.DB) error {
	return c.validate()
}

func (c *Category) validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if len(c.Name) > 255 {
		return errors.New("name must be 255 characters or less")
	}
	return nil
}

// NamesByKeywords is equivalent to the Rails scope
func CategoryNamesByKeywords(db *gorm.DB, keywords []string) *gorm.DB {
	if len(keywords) == 0 {
		return db.Where("1 = 0") // Equivalent to ActiveRecord's `none`
	}

	query := db
	for _, keyword := range keywords {
		likeKeyword := "%" + keyword + "%"
		query = query.Where("name LIKE ?", likeKeyword)
	}
	return query
}