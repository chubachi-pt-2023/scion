package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID                     uint      `gorm:"primaryKey"`
	ProductID              uint      `gorm:"not null"`
	DiscontinuedProductID  uint      `gorm:"index;not null"`
	UserID                 uint      `gorm:"not null"`
	Comment                string    `gorm:"type:text;not null"`
	Anonymous              bool      `gorm:"default:false;not null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time

	User                   *User     `gorm:"foreignKey:UserID"`
	Product                *Product  `gorm:"foreignKey:ProductID"`
	DiscontinuedProduct    *Product  `gorm:"foreignKey:DiscontinuedProductID"`
}

func (Review) TableName() string {
	return "reviews"
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	var count int64
	err := tx.Model(&Review{}).
		Where("product_id = ? AND discontinued_product_id = ? AND user_id = ?", r.ProductID, r.DiscontinuedProductID, r.UserID).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("review already exists for this product, discontinued product, and user combination")
	}
	return nil
}

func (r *Review) Validate() error {
	if r.Comment == "" {
		return errors.New("comment cannot be empty")
	}
	return nil
}