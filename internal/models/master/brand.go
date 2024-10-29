package master

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Brand struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	NameJa    *string   `gorm:"column:name_ja;type:varchar(255);uniqueIndex"`
	NameEn    *string   `gorm:"column:name_en;type:varchar(255);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Brand) TableName() string {
	return "master_brands"
}

func (b *Brand) BeforeCreate(tx *gorm.DB) error {
	return b.validate()
}

func (b *Brand) BeforeUpdate(tx *gorm.DB) error {
	return b.validate()
}

func (b *Brand) validate() error {
	if b.Name == "" {
		return errors.New("name is required")
	}
	if len(b.Name) > 255 {
		return errors.New("name must be 255 characters or less")
	}
	if b.NameJa != nil && len(*b.NameJa) > 255 {
		return errors.New("name_ja must be 255 characters or less")
	}
	if b.NameEn != nil && len(*b.NameEn) > 255 {
		return errors.New("name_en must be 255 characters or less")
	}
	return nil
}

// NamesByKeywords is equivalent to the Rails scope
func NamesByKeywords(db *gorm.DB, keywords []string) *gorm.DB {
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