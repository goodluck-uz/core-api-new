package postgres

import (
	"context"
	"fmt"

	"github.com/goodluck-uz/core-api/config"
	"github.com/goodluck-uz/core-api/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	category storage.CategoryRepoI
}

var logPath = "storage/postgres/postgresql.go"

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, fmt.Errorf(logPath, " NewConnectPostgresql: ParseConfig: %w", err)
	}
	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf(logPath, " NewConnectPostgresql: ConnectConfig: %w", err)
	}

	return &Store{
		db: pgpool,
	}, nil

}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}
	return s.category
}
