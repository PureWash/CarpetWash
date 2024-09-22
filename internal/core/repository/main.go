package repository

import (
	"context"
	"fmt"

	"carpet/internal/configs"

	"carpet/internal/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	DB  *pgxpool.Pool
	log logger.ILogger
	cfg configs.Config
}

func NewStore(ctx context.Context, log logger.ILogger, cnf configs.Config) (*Store, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cnf.PostgresUser,
		cnf.PostgresPassword,
		cnf.PostrgresHost,
		cnf.PostrgresPort,
		cnf.PostgresDatabase,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Error("this error is parse url -> can not parsing", logger.Error(err))
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Error("this error ie new create config with pool", logger.Error(err))
		return nil, err
	}
	return &Store{
		DB:  pool,
		log: log,
		cfg: cnf,
	}, nil
}
