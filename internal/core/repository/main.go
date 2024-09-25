package repository

import (
	"carpet/internal/configs"
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/pkg/logger"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	DB      *pgxpool.Pool
	log     logger.ILogger
	cfg     configs.Config
	Queries *sqlc.Queries
}

func NewStore(ctx context.Context, log logger.ILogger, cnf configs.Config) (*Store, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cnf.PostgresUser,
		cnf.PostgresPassword,
		cnf.PostrgresHost,
		cnf.PostrgresPort,
		cnf.PostgresDatabase,
	)
	// fmt.Println(url)

	pool, err := pgxpool.Connect(ctx, url)
	if err != nil {
		log.Error("Error creating connection pool", logger.Error(err))
		return nil, err
	}

	queries := sqlc.New(pool)

	return &Store{
		DB:      pool,
		log:     log,
		cfg:     cnf,
		Queries: queries,
	}, nil
}
