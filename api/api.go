package api

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/goodluck-uz/core-api/api/docs"
	"github.com/goodluck-uz/core-api/api/handler"
	"github.com/goodluck-uz/core-api/config"
	"github.com/goodluck-uz/core-api/pkg/logger"
	"github.com/goodluck-uz/core-api/storage"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func NewApi(app *fiber.App, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)
	// Category
	app.Post("/categories", handler.CreateCategory)

	// User
	app.Post("/users", handler.CreateUser)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

}
