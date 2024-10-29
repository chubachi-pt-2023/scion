package models

import (
	"atomono-api/internal/models/master"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                 uint           `gorm:"primaryKey"`
	Name               string         `gorm:"type:varchar(255)"`
	CompanyID          *uint          `gorm:"index"`
	BrandID            *uint          `gorm:"index"`
	CategoryID         *uint          `gorm:"index"`
	CountryID          *uint          `gorm:"index"`
	ImageURL           string         `gorm:"type:text"`
	OriginImageURL     string         `gorm:"type:text"`
	Price              *uint          
	Description        string         `gorm:"type:text"`
	Composition        string         `gorm:"type:text"`
	Capacity           string         `gorm:"type:varchar(255)"`
	Size               string         `gorm:"type:varchar(255)"`
	ProductStatus      ProductStatus  `gorm:"type:integer;default:1;not null"`
	DiscontinuedOn     *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Company            *master.Company       `gorm:"foreignKey:CompanyID"`
	Brand              *master.Brand         `gorm:"foreignKey:BrandID"`
	Category           *master.Category      `gorm:"foreignKey:CategoryID"`
	Country            *master.Country       `gorm:"foreignKey:CountryID"`
	ReplacesProducts   []ReplacesProduct     `gorm:"foreignKey:OldProductID"`
	NewProducts        []Product             `gorm:"many2many:replaces_products;joinForeignKey:OldProductID;joinReferences:NewProductID"`
	UsersProducts      []UsersProduct
	Users              []User                `gorm:"many2many:users_products"`
	Reviews            []Review              `gorm:"foreignKey:ProductID"`
	DiscontinuedReviews[]Review              `gorm:"foreignKey:DiscontinuedProductID"`
}

type ProductStatus int

const (
	ProductStatusDiscontinued ProductStatus = iota
	ProductStatusCurrent
	ProductStatusOtherwise
)

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ProductStatus == 0 {
		p.ProductStatus = ProductStatusCurrent
	}
	return nil
}

func (p *Product) Validate() error {
	if len(p.Name) > 255 {
		return errors.New("product name must be 255 characters or less")
	}
	if len(p.ImageURL) > 50000 || len(p.OriginImageURL) > 50000 || len(p.Description) > 50000 || len(p.Composition) > 50000 {
		return errors.New("product image_url, origin_image_url, description, and composition must be 50000 characters or less")
	}
	if p.Price != nil && *p.Price < 0 {
		return errors.New("price must be greater than or equal to 0")
	}
	return nil
}

func (Product) TableName() string {
	return "products"
}

// NamesLikeBy is equivalent to the Rails scope
func NamesLikeBy(db *gorm.DB, keywords []string) *gorm.DB {
	if len(keywords) == 0 {
		return db.Where("1 = 0") // Equivalent to ActiveRecord's `none`
	}

	query := db
	for _, keyword := range keywords {
		query = query.Or("name LIKE ?", "%"+keyword+"%")
	}
	return query
}