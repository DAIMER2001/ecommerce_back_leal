package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitDB(cfg Config) (*pgxpool.Pool, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUsername,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)

	db, err := pgxpool.Connect(context.Background(), connectionString)

	return db, err
}
