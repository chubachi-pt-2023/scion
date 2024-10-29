package middleware

import (
	"atomono-api/internal/services"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(userService *services.UserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authorization header"})
			}

			bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
			if bearerToken == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			}

			// ここでトークンを検証し、ユーザーIDを取得する処理を実装
			// この例では簡略化のため、トークンをユーザーIDとして扱います
			userID, err := validateToken(bearerToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			user, err := userService.GetUserByID(userID)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not found"})
			}

			c.Set("user_id", user.ID)
			return next(c)
		}
	}
}

func validateToken(token string) (uint, error) {
	// ここで実際のトークン検証ロジックを実装
	// この例では簡略化のため、トークンを直接ユーザーIDとして扱います
	// 実際の実装では、JWTなどの適切な認証メカニズムを使用してください
	userID := uint(1) // 仮のユーザーID
	return userID, nil
}