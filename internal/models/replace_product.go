package models

import (
	"time"

	"gorm.io/gorm"
)

type ReplacesProduct struct {
	ID            uint         `gorm:"primaryKey"`
	OldProductID  uint         `gorm:"index;not null"`
	NewProductID  uint         `gorm:"index;not null"`
	SourceKind    SourceKind   `gorm:"type:integer;default:0;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	OldProduct    *Product     `gorm:"foreignKey:OldProductID"`
	NewProduct    *Product     `gorm:"foreignKey:NewProductID"`
}

type SourceKind int

const (
	SourceKindMaker SourceKind = iota
	SourceKindEC
	SourceKindUser
	SourceKindCosineSimilarity
)

func (ReplacesProduct) TableName() string {
	return "replaces_products"
}

func (rp *ReplacesProduct) BeforeCreate(tx *gorm.DB) error {
	if rp.SourceKind == 0 {
		rp.SourceKind = SourceKindMaker
	}
	return nil
}