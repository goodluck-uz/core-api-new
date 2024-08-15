package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	Environment string // debug, test, release

	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
	FileName         string

	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {
	cfg := Config{}
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.ServerHost = "localhost"
	cfg.ServerPort = "9000"

	cfg.PostgresHost = os.Getenv("POSTGRES_HOST")
	cfg.PostgresUser = os.Getenv("POSTGRES_USER")
	cfg.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
	cfg.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	cfg.PostgresPort = os.Getenv("POSTGRES_PORT")
	cfg.FileName = "../api-logs/app.log"
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	return cfg
}
