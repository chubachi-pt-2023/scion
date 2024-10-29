package services

import (
	"atomono-api/internal/models"
	"atomono-api/internal/repositories"
	"errors"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *UserService) GetUserByToken(token string) (*models.User, error) {
	// ここでトークンを検証し、対応するユーザーを返す実装を行います
	// 例：
	user, err := s.userRepo.FindByToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	return user, nil
}