package handlers

import (
	"atomono-api/internal/models"
	"atomono-api/internal/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService *services.ProductService
	userService    *services.UserService
}

func NewProductHandler(ps *services.ProductService, us *services.UserService) *ProductHandler {
	return &ProductHandler{
		productService: ps,
		userService:    us,
	}
}

func (h *ProductHandler) Index(c echo.Context) error {
    products, err := h.productService.GetProducts(100) // Limit to 100 products
    if err != nil {
        log.Printf("Error fetching products: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
    }

    if len(products) == 0 {
        log.Println("No products found")
        return c.JSON(http.StatusOK, []interface{}{})
    }

    log.Printf("Found %d products", len(products))
    return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	return c.JSON(http.StatusOK, formatResponse(product))
}

func (h *ProductHandler) Search(c echo.Context) error {
	var searchRequest struct {
		Query []string `json:"query"`
	}

	if err := json.NewDecoder(c.Request().Body).Decode(&searchRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	products, err := h.productService.SearchProducts(searchRequest.Query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search products"})
	}

	return c.JSON(http.StatusOK, formatSearchResponse(products))
}

// Helper functions to format the response
func formatSearchResponse(products []models.Product) []map[string]interface{} {
	// Implement the formatting logic here
	// This should match the format used in your Ruby implementation
	return nil
}

func formatResponse(product *models.Product) map[string]interface{} {
	// Implement the formatting logic here
	// This should match the format used in your Ruby implementation
	return nil
}