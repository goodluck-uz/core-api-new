package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goodluck-uz/core-api/api/models"
)

var logPath1 = "api/handler/user.go"

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param input body models.CreateUserRequest true "Create user"
// @Success 200 {object} models.User
// @Router /users [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return h.handlerResponse(c, logPath1+"create.parse", fiber.StatusBadRequest, err.Error())
	}

	user, err := h.storages.User().Create(c.Context(), &req)
	if err != nil {
		return h.handlerResponse(c, logPath1+" storage.user.create", fiber.StatusInternalServerError, err.Error())
	}

	return h.handlerResponse(c, "create user", fiber.StatusCreated, user)

}
