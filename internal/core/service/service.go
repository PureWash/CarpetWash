package service

import (
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/pkg/logger"
)

type Service struct {
	storage sqlc.Querier
	log     logger.ILogger
}

func NewService(storage sqlc.Querier, log logger.ILogger) *Service {
	return &Service{
		storage: storage,
		log: log,
	}
}