package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitDB(cfg Config) (*pgxpool.Pool, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d",
		cfg.PostgresUsername,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
		cfg.PostgresMaxConns,
	)

	db, err := pgxpool.Connect(context.Background(), connectionString)

	return db, err
}
