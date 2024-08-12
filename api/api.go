package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goodluck-uz/core-api/api/handler"
	_ "github.com/goodluck-uz/core-api/cmd/docs"
	"github.com/goodluck-uz/core-api/config"
	"github.com/goodluck-uz/core-api/pkg/logger"
	"github.com/goodluck-uz/core-api/storage"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func NewApi(app *fiber.App, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)

	app.Post("/categories", handler.CreateCategory)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

}
