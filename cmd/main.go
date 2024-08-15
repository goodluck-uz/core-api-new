package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/goodluck-uz/core-api/api"
	"github.com/goodluck-uz/core-api/config"
	"github.com/goodluck-uz/core-api/pkg/logger"
	postgresql "github.com/goodluck-uz/core-api/storage/postgres"
	"go.uber.org/zap"
)

func main() {

	cfg := config.Load()

	// Initialize logger level based on environment
	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
	default:
		*loggerLevel = logger.LevelInfo
	}
	log := logger.NewLogger("app", *loggerLevel, cfg.FileName)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()
	// Connect to PostgreSQL
	store, err := postgresql.NewConnectPostgresql(&cfg)
	if err != nil {
		log.Panic("Error connecting to PostgreSQL", zap.Error(err)) // Panic if connection fails
		return
	}
	defer store.CloseDB()

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		// Add Fiber configuration options here if needed
		Prefork: false, // Use prefork if necessary
	})

	// Set up middleware (equivalent to Gin's Recovery and Logger)
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				log.Error("Panic recovered: ", logger.Any("error", err))
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
			}
		}()
		return c.Next()
	})

	// You can add a custom logger middleware if needed
	// app.Use(loggerMiddleware)

	// Set up API routes
	api.NewApi(app, &cfg, store, log)

	// Start Fiber server
	serverAddress := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	fmt.Println("Server running on port", serverAddress)

	if err := app.Listen(serverAddress); err != nil {
		log.Panic("Error starting server: ", logger.Error(err))
		return
	}
}
