package master

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const DevelopmentDataID = 1 // 開発用データに紐づく企業ID

type Company struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	NameJa    *string   `gorm:"column:name_ja;type:varchar(255);uniqueIndex"`
	NameEn    *string   `gorm:"column:name_en;type:varchar(255);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Company) TableName() string {
	return "master_companies"
}

func (c *Company) BeforeCreate(tx *gorm.DB) error {
	return c.validate()
}

func (c *Company) BeforeUpdate(tx *gorm.DB) error {
	return c.validate()
}

func (c *Company) validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if len(c.Name) > 255 {
		return errors.New("name must be 255 characters or less")
	}
	if c.NameJa != nil && len(*c.NameJa) > 255 {
		return errors.New("name_ja must be 255 characters or less")
	}
	if c.NameEn != nil && len(*c.NameEn) > 255 {
		return errors.New("name_en must be 255 characters or less")
	}
	return nil
}

// NamesLikeBy is equivalent to the Rails scope
func NamesLikeBy(db *gorm.DB, keywords []string) *gorm.DB {
	if len(keywords) == 0 {
		return db.Where("1 = 0") // Equivalent to ActiveRecord's `none`
	}

	query := db
	for _, keyword := range keywords {
		likeKeyword := "%" + keyword + "%"
		query = query.Where("name LIKE ? OR name_ja LIKE ? OR name_en LIKE ?", likeKeyword, likeKeyword, likeKeyword)
	}
	return query
}