package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goodluck-uz/core-api/api/models"
)

var logPath = "api/handler/category.go"

// CreateCategory godoc
// @Summary Create category
// @Description Create category
// @Tags categories
// @Accept json
// @Produce json
// @Param input body models.CreateCategoryRequest true "Create category"
// @Success 200 {object} models.Category
// @Router /categories [post]
func (h *Handler) CreateCategory(c *fiber.Ctx) error {
	var req models.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return h.handlerResponse(c, logPath+"create.parse", fiber.StatusBadRequest, err.Error())
	}

	category, err := h.storages.Category().Create(c.Context(), &req)
	if err != nil {
		return h.handlerResponse(c, logPath+" storage.category.create", fiber.StatusInternalServerError, err.Error())
	}

	return h.handlerResponse(c, "create category", fiber.StatusCreated, category)

}
