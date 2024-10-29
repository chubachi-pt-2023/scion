package services

import (
	"atomono-api/internal/models"
	"atomono-api/internal/repositories"
	"errors"
)

var ErrReviewNotFound = errors.New("review not found")

type ReviewService struct {
	reviewRepo *repositories.ReviewRepository
}

func NewReviewService(rr *repositories.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRepo: rr}
}

func (s *ReviewService) GetReviewsByDiscontinuedProductID(productID uint) ([]models.Review, error) {
	return s.reviewRepo.FindByDiscontinuedProductID(productID)
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.reviewRepo.Create(review)
}

func (s *ReviewService) DeleteReview(reviewID, productID, userID uint) error {
	review, err := s.reviewRepo.FindByID(reviewID)
	if err != nil {
		return ErrReviewNotFound
	}

	if review.DiscontinuedProductID != productID || review.UserID != userID {
		return ErrReviewNotFound
	}

	return s.reviewRepo.Delete(review)
}