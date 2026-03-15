package database

import (
	"context"
	"fmt"
	"time"

	"github.com/djwhocodes/go-grpc-microservices/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
