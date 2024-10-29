package services

import (
	"atomono-api/internal/models"
	"atomono-api/internal/repositories"
)

type ReplacesProductService struct {
    repo *repositories.ReplacesProductRepository
}

func NewReplacesProductService(repo *repositories.ReplacesProductRepository) *ReplacesProductService {
    return &ReplacesProductService{repo: repo}
}

func (s *ReplacesProductService) CreateReplacesProduct(oldProductID, newProductID uint, sourceKind models.SourceKind) error {
    rp := &models.ReplacesProduct{
        OldProductID: oldProductID,
        NewProductID: newProductID,
        SourceKind:   sourceKind,
    }
    return s.repo.Create(rp)
}

func (s *ReplacesProductService) GetReplacementsForProduct(productID uint) ([]models.ReplacesProduct, error) {
    return s.repo.FindByOldProductID(productID)
}

// その他必要なメソッドを追加