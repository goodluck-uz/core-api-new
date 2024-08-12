package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/goodluck-uz/core-api/config"
	"github.com/goodluck-uz/core-api/pkg/logger"
	"github.com/goodluck-uz/core-api/storage"
)

type Handler struct {
	cfg      *config.Config
	logger   logger.LoggerI
	storages storage.StorageI
}
type Response struct {
	Status      int
	Description string
	Data        interface{}
}

func NewHandler(cfg *config.Config, store storage.StorageI, logger logger.LoggerI) *Handler {
	return &Handler{
		cfg:      cfg,
		logger:   logger,
		storages: store,
	}
}

func (h *Handler) handlerResponse(c *fiber.Ctx, path string, code int, message interface{}) error {
	response := Response{
		Status:      code,
		Description: path,
		Data:        message,
	}

	switch {
	case code < 300:
		h.logger.Info(path, logger.Any("info", response))
	case code >= 400:
		h.logger.Error(path, logger.Any("error", response))
	}

	return c.Status(code).JSON(response)
}

func (h *Handler) getOffsetQuery(offset string) (int, error) {
	if len(offset) <= 0 {
		return h.cfg.DefaultOffset, nil
	}

	return strconv.Atoi(offset)
}

func (h *Handler) getLimitQuery(limit string) (int, error) {

	if len(limit) <= 0 {
		return h.cfg.DefaultLimit, nil
	}

	return strconv.Atoi(limit)
}
