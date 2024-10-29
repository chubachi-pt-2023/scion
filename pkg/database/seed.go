package database

import (
	"atomono-api/internal/models"
	"atomono-api/internal/models/master"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	log.Println("Seeding database...")

	// Countries
	countries := []master.Country{
		{Name: "Japan"},
		{Name: "United States"},
		{Name: "France"},
		{Name: "South Korea"},
	}
	if err := db.Create(&countries).Error; err != nil {
		return err
	}
	ptr := func(s string) *string {
		return &s
	}


	// Companies
	companies := []master.Company{
		{Name: "Shiseido", NameJa: ptr("資生堂"), NameEn: ptr("Shiseido")},
		{Name: "L'Oréal", NameJa: ptr("ロレアル"), NameEn: ptr("L'Oreal")},
		{Name: "Estée Lauder", NameJa: ptr("エスティローダー"), NameEn: ptr("Estee Lauder")},
		{Name: "AmorePacific", NameJa: ptr("アモーレパシフィック"), NameEn: ptr("AmorePacific")},
	}
	if err := db.Create(&companies).Error; err != nil {
		return err
	}

	// Brands
	brands := []master.Brand{
		{Name: "SHISEIDO", NameJa: ptr("シセイドウ"), NameEn: ptr("SHISEIDO")},
		{Name: "Lancôme", NameJa: ptr("ランコム"), NameEn: ptr("Lancome")},
		{Name: "Clinique", NameJa: ptr("クリニーク"), NameEn: ptr("Clinique")},
		{Name: "Laneige", NameJa: ptr("ラネージュ"), NameEn: ptr("Laneige")},
	}
	if err := db.Create(&brands).Error; err != nil {
		return err
	}

	// Categories
	categories := []master.Category{
		{Name: "Skincare"},
		{Name: "Makeup"},
		{Name: "Fragrance"},
		{Name: "Haircare"},
	}
	if err := db.Create(&categories).Error; err != nil {
		return err
	}

	// Products
	products := []models.Product{
		{
			Name:           "Ultimune Power Infusing Concentrate",
			CompanyID:      &companies[0].ID,
			BrandID:        &brands[0].ID,
			CategoryID:     &categories[0].ID,
			CountryID:      &countries[0].ID,
			ImageURL:       "https://example.com/ultimune.jpg",
			Price:          func() *uint { v := uint(13000); return &v }(),
			Description:    "Strengthens skin's inner defenses and helps restore radiance.",
			Composition:    "Water, Glycerin, Dimethicone, Butylene Glycol...",
			Capacity:       "50ml",
			ProductStatus:  models.ProductStatusCurrent,
		},
		{
			Name:           "Advanced Génifique Youth Activating Serum",
			CompanyID:      &companies[1].ID,
			BrandID:        &brands[1].ID,
			CategoryID:     &categories[0].ID,
			CountryID:      &countries[2].ID,
			ImageURL:       "https://example.com/genifique.jpg",
			Price:          func() *uint { v := uint(14000); return &v }(),
			Description:    "Targets 10 key signs of aging in just 7 days.",
			Composition:    "Aqua / Water, Bifida Ferment Lysate, Glycerin, Alcohol Denat...",
			Capacity:       "30ml",
			ProductStatus:  models.ProductStatusCurrent,
		},
		{
			Name:           "Moisture Surge 100H Auto-Replenishing Hydrator",
			CompanyID:      &companies[2].ID,
			BrandID:        &brands[2].ID,
			CategoryID:     &categories[0].ID,
			CountryID:      &countries[1].ID,
			ImageURL:       "https://example.com/moisturesurge.jpg",
			Price:          func() *uint { v := uint(5500); return &v }(),
			Description:    "Oil-free gel-cream that provides 100 hours of hydration.",
			Composition:    "Water, Dimethicone, Butylene Glycol, Glycerin...",
			Capacity:       "50ml",
			ProductStatus:  models.ProductStatusCurrent,
		},
		{
			Name:           "Water Sleeping Mask",
			CompanyID:      &companies[3].ID,
			BrandID:        &brands[3].ID,
			CategoryID:     &categories[0].ID,
			CountryID:      &countries[3].ID,
			ImageURL:       "https://example.com/sleepingmask.jpg",
			Price:          func() *uint { v := uint(3900); return &v }(),
			Description:    "Overnight mask that purifies and hydrates skin while you sleep.",
			Composition:    "Water, Butylene Glycol, Cyclopentasiloxane, Glycerin...",
			Capacity:       "70ml",
			ProductStatus:  models.ProductStatusCurrent,
		},
	}
	if err := db.Create(&products).Error; err != nil {
		return err
	}

	// Users
	users := []models.User{
		{
			Email:    "user1@example.com",
			// Password: "hashed_password_here", // 実際にはハッシュ化したパスワードを保存します
		},
		{
			Email:    "user2@example.com",
			// Password: "hashed_password_here",
		},
	}
	if err := db.Create(&users).Error; err != nil {
		return err
	}

	// Reviews
	reviews := []models.Review{
		{
			ProductID:  products[0].ID,
			UserID:     users[0].ID,
			Comment:    "Great product! My skin feels amazing.",
			Anonymous:  false,
		},
		{
			ProductID:  products[1].ID,
			UserID:     users[1].ID,
			Comment:    "I've been using this for a month and can see the difference.",
			Anonymous:  true,
		},
	}
	if err := db.Create(&reviews).Error; err != nil {
		return err
	}

	log.Println("Seeding completed successfully")
	return nil
}