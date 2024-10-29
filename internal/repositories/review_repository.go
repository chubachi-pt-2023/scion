package repositories

import (
	"atomono-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) FindByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("review not found")
		}
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) FindByDiscontinuedProductID(productID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("discontinued_product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) Delete(review *models.Review) error {
	return r.db.Delete(review).Error
}

func (r *ReviewRepository) Update(review *models.Review) error {
	return r.db.Save(review).Error
}