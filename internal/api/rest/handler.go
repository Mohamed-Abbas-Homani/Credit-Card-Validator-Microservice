package rest

import (
	"net/http"

	"credit-card-validator/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	validator *service.Validator
	logger    *logrus.Logger
}

type ValidateRequest struct {
	CardNumber string `json:"card_number" validate:"required"`
}

func NewHandler(validator *service.Validator, logger *logrus.Logger) *Handler {
	return &Handler{
		validator: validator,
		logger:    logger,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.POST("/validate", h.ValidateCard)
}

func (h *Handler) ValidateCard(c echo.Context) error {
	var req ValidateRequest
	if err := c.Bind(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	if req.CardNumber == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Card number is required",
		})
	}

	result := h.validator.ValidateCard(req.CardNumber)
	return c.JSON(http.StatusOK, result)
}
