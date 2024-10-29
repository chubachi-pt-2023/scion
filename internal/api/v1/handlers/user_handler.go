package handlers

import (
	"atomono-api/internal/models"
	"atomono-api/internal/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
    userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
    return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) GetUser(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := h.userRepo.FindByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }
    return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
    user := new(models.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }
    if err := h.userRepo.Create(user); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
    }
    return c.JSON(http.StatusCreated, user)
}

// 他のハンドラーメソッド（UpdateUser, DeleteUser など）も同様に実装