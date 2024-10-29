package middleware

import (
	"atomono-api/internal/services"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth(userService *services.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authorization token"})
			}

			// "Bearer " プレフィックスを削除
			token = strings.TrimPrefix(token, "Bearer ")

			user, err := userService.GetUserByToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			c.Set("user", user)
			return next(c)
		}
	}
}