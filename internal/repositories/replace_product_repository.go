package repositories

import (
	"atomono-api/internal/models"

	"gorm.io/gorm"
)

type ReplacesProductRepository struct {
    db *gorm.DB
}

func NewReplacesProductRepository(db *gorm.DB) *ReplacesProductRepository {
    return &ReplacesProductRepository{db: db}
}

func (r *ReplacesProductRepository) Create(rp *models.ReplacesProduct) error {
    return r.db.Create(rp).Error
}

func (r *ReplacesProductRepository) FindByOldProductID(oldProductID uint) ([]models.ReplacesProduct, error) {
    var replacesProducts []models.ReplacesProduct
    err := r.db.Where("old_product_id = ?", oldProductID).Find(&replacesProducts).Error
    return replacesProducts, err
}

func (r *ReplacesProductRepository) FindByNewProductID(newProductID uint) ([]models.ReplacesProduct, error) {
    var replacesProducts []models.ReplacesProduct
    err := r.db.Where("new_product_id = ?", newProductID).Find(&replacesProducts).Error
    return replacesProducts, err
}

// その他必要なメソッドを追加