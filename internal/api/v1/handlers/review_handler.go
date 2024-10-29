package handlers

import (
	"atomono-api/internal/models"
	"atomono-api/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	reviewService  *services.ReviewService
	productService *services.ProductService
	userService    *services.UserService
}

func NewReviewHandler(rs *services.ReviewService, ps *services.ProductService, us *services.UserService) *ReviewHandler {
	return &ReviewHandler{
		reviewService:  rs,
		productService: ps,
		userService:    us,
	}
}

func (h *ReviewHandler) GetReviews(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	discontinuedProduct, err := h.productService.GetDiscontinuedProductByID(uint(productID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Discontinued product not found"})
	}

	reviews, err := h.reviewService.GetReviewsByDiscontinuedProductID(discontinuedProduct.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch reviews"})
	}

	return c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	var reviewRequest struct {
		Comment   string `json:"comment"`
		Anonymous bool   `json:"anonymous"`
	}

	if err := c.Bind(&reviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	userID := c.Get("user_id").(uint) // Assuming middleware sets this

	discontinuedProduct, err := h.productService.GetDiscontinuedProductByID(uint(productID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Discontinued product not found"})
	}

	review := &models.Review{
		ProductID:             discontinuedProduct.ID,
		DiscontinuedProductID: discontinuedProduct.ID,
		UserID:                userID,
		Comment:               reviewRequest.Comment,
		Anonymous:             reviewRequest.Anonymous,
	}

	if err := h.reviewService.CreateReview(review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create review"})
	}

	return c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid review ID"})
	}

	userID := c.Get("user_id").(uint) // Assuming middleware sets this

	err = h.reviewService.DeleteReview(uint(reviewID), uint(productID), userID)
	if err != nil {
		if err == services.ErrReviewNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete review"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":               []string{"deleted_discontinued_product_review"},
		"reviewId":               reviewID,
		"discontinuedProductId":  productID,
		"productId":              productID,
	})
}