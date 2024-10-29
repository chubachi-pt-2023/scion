package repositories

import (
	"atomono-api/internal/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindDiscontinuedByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND product_status = ?", id, models.ProductStatusDiscontinued).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("discontinued product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) FindAll(limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Limit(limit).Find(&products).Error
	return products, err
}

func (r *ProductRepository) SearchByKeywords(keywords []string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("name LIKE ?", "%"+strings.Join(keywords, "%")+"%").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Limit(limit int) *gorm.DB {
	return r.db.Limit(limit)
}