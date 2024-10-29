package main

import (
	"atomono-api/internal/api/v1/handlers"
	"atomono-api/internal/api/v1/routes"
	"atomono-api/internal/repositories"
	"atomono-api/internal/services"
	"atomono-api/pkg/config"
	"atomono-api/pkg/database"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
    config.LoadConfig()
	db := database.Init()

	// マイグレーションの実行
	database.Migrate(db)
	 // シードデータの追加
	 if os.Getenv("ATOMONO_ENV") == "development" {
		if err := database.Seed(db); err != nil {
			log.Fatalf("Error seeding database: %v", err)
		}
	}
	
	// Repositories
	replacesProductRepo := repositories.NewReplacesProductRepository(db)
	reviewRepo := repositories.NewReviewRepository(db)
	productRepo := repositories.NewProductRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Services
	replacesProductService := services.NewReplacesProductService(replacesProductRepo)
	reviewService := services.NewReviewService(reviewRepo)
	productService := services.NewProductService(productRepo)
	userService := services.NewUserService(userRepo)

	// Handlers
	replacesProductHandler := handlers.NewReplacesProductHandler(replacesProductService)
	reviewHandler := handlers.NewReviewHandler(reviewService, productService, userService)
	productHandler := handlers.NewProductHandler(productService, userService)

	// Echo instance
	e := echo.New()

	// Setup routes
	routes.SetupRoutes(e, replacesProductHandler, productHandler, reviewHandler, userService)

	// Start server
	// e.Logger.Fatal(e.Start(":8080"))
    e.Logger.Fatal(e.Start(":" + config.AppConfig.ServerPort))
}