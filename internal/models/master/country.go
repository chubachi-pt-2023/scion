package master

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Country) TableName() string {
	return "master_countries"
}

func (c *Country) BeforeCreate(tx *gorm.DB) error {
	return c.validate()
}

func (c *Country) BeforeUpdate(tx *gorm.DB) error {
	return c.validate()
}

func (c *Country) validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if len(c.Name) > 255 {
		return errors.New("name must be 255 characters or less")
	}
	return nil
}