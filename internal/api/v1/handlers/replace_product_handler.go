package handlers

import (
	"atomono-api/internal/models"
	"atomono-api/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReplacesProductHandler struct {
    service *services.ReplacesProductService
}

func NewReplacesProductHandler(service *services.ReplacesProductService) *ReplacesProductHandler {
    return &ReplacesProductHandler{service: service}
}

func (h *ReplacesProductHandler) CreateReplacesProduct(c echo.Context) error {
    var req struct {
        OldProductID uint              `json:"old_product_id"`
        NewProductID uint              `json:"new_product_id"`
        SourceKind   models.SourceKind `json:"source_kind"`
    }

    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    err := h.service.CreateReplacesProduct(req.OldProductID, req.NewProductID, req.SourceKind)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create replaces product"})
    }

    return c.JSON(http.StatusCreated, map[string]string{"message": "Replaces product created successfully"})
}

func (h *ReplacesProductHandler) GetReplacements(c echo.Context) error {
    productID, err := strconv.ParseUint(c.Param("productID"), 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
    }

    replacements, err := h.service.GetReplacementsForProduct(uint(productID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get replacements"})
    }

    return c.JSON(http.StatusOK, replacements)
}

// その他必要なハンドラーメソッドを追加