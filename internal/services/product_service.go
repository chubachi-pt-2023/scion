package services

import (
	"atomono-api/internal/models"
	"atomono-api/internal/repositories"
	"log"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(pr *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: pr}
}

func (s *ProductService) GetProducts(limit int) ([]models.Product, error) {
    var products []models.Product
    result := s.productRepo.Limit(limit).Find(&products)
    if result.Error != nil {
        log.Printf("Error fetching products from database: %v", result.Error)
        return nil, result.Error
    }
    log.Printf("Fetched %d products from database", len(products))
    return products, nil
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		log.Printf("Error fetching product with ID %d: %v", id, err)
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetDiscontinuedProductByID(productID uint) (*models.Product, error) {
	product, err := s.productRepo.FindDiscontinuedByID(productID)
	if err != nil {
		log.Printf("Error fetching discontinued product with ID %d: %v", productID, err)
		return nil, err
	}
	return product, nil
}

func (s *ProductService) SearchProducts(keywords []string) ([]models.Product, error) {
	products, err := s.productRepo.SearchByKeywords(keywords)
	if err != nil {
		log.Printf("Error searching products with keywords %v: %v", keywords, err)
		return nil, err
	}
	return products, nil
}