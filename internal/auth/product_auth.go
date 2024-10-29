package auth

import (
	"atomono-api/internal/models"
)

type ProductAuth struct {
	User *models.User
}

func NewProductAuth(user *models.User) *ProductAuth {
	return &ProductAuth{User: user}
}

func (a *ProductAuth) CanIndex() bool {
	return a.User != nil
}

func (a *ProductAuth) CanCreate() bool {
	return a.User != nil
}

func (a *ProductAuth) CanUpdate(userID uint) bool {
	return a.User != nil && a.User.ID == userID
}

func (a *ProductAuth) CanDelete(userID uint) bool {
	return a.User != nil && a.User.ID == userID
}