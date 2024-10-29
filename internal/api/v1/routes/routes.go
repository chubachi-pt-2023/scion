package routes

import (
	"atomono-api/internal/api/v1/handlers"
	"atomono-api/internal/services"
	"atomono-api/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(
	e *echo.Echo,
	rph *handlers.ReplacesProductHandler,
	ph *handlers.ProductHandler,
	rh *handlers.ReviewHandler,
	userService *services.UserService,
) {
	v1 := e.Group("/api/v1")

	// Public routes
	v1.GET("/products", ph.Index)
	v1.GET("/products/:id", ph.Show)
	v1.POST("/products/search", ph.Search)

	// ReplacesProduct routes
	v1.POST("/replaces-products", rph.CreateReplacesProduct)
	v1.GET("/products/:productID/replacements", rph.GetReplacements)

	// Product reviews routes (with authentication)
	products := v1.Group("/products/:productId")
	products.Use(middleware.Auth(userService))
	products.GET("/reviews", rh.GetReviews)
	products.POST("/reviews", rh.CreateReview)
	products.DELETE("/reviews/:id", rh.DeleteReview)

	// User products routes (with authentication)
	userProducts := v1.Group("/users/:userId/products")
	userProducts.Use(middleware.Auth(userService))
	// Add user product routes here, e.g.:
	// userProducts.GET("", ph.GetUserProducts)
	// userProducts.POST("", ph.CreateUserProduct)
}