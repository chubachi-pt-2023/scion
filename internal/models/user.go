package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex"`
	UID       string    `gorm:"type:varchar(255);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UsersProducts []UsersProduct `gorm:"foreignKey:UserID"`
	Products      []Product      `gorm:"many2many:users_products"`
	Profile       *UsersProfile  `gorm:"foreignKey:UserID"`
	Favorites     []UsersFavorite `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeValidate(tx *gorm.DB) error {
	u.Name = strings.TrimSpace(u.Name)
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	if u.Profile == nil {
		profile := &UsersProfile{
			UserID:    u.ID,
			SkinType:  UsersProfileDefaultValue,
			SkinColor: UsersProfileDefaultValue,
			HairType:  UsersProfileDefaultValue,
		}
		if err := tx.Create(profile).Error; err != nil {
			return err
		}
		u.Profile = profile
	}
	return nil
}

func (u *User) Validate() error {
	if len(u.Name) > 255 {
		return errors.New("name must be 255 characters or less")
	}
	if len(u.Email) > 255 {
		return errors.New("email must be 255 characters or less")
	}
	if len(u.UID) > 255 {
		return errors.New("uid must be 255 characters or less")
	}
	return nil
}

// Constants for UsersProfile
const UsersProfileDefaultValue = "default_value"

type UsersProfile struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"uniqueIndex"`
	SkinType  string
	SkinColor string
	HairType  string
	CreatedAt time.Time
	UpdatedAt time.Time

	User *User `gorm:"foreignKey:UserID"`
}

func (UsersProfile) TableName() string {
	return "users_profiles"
}

type UsersProduct struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time

	User    *User    `gorm:"foreignKey:UserID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}

func (UsersProduct) TableName() string {
	return "users_products"
}

type UsersFavorite struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time

	User    *User    `gorm:"foreignKey:UserID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}

func (UsersFavorite) TableName() string {
	return "users_favorites"
}