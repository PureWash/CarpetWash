package main

import (
	"carpet/internal/configs"
	"carpet/internal/core/repository"
	"carpet/internal/pkg/grpcConn"
	"carpet/internal/pkg/logger"
	"context"
	"net"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()
	var loggerLevel string
	switch cfg.Environment {
	case logger.LevelDebug:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case configs.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	potgresStore, err := repository.NewStore(context.Background(), log, cfg)
	if err != nil {
		log.Error("error while connecting to postgres", logger.Error(err))
		return
	}
	grpcServer := grpcConn.ConnGRPC(potgresStore.Queries, log)

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Error("error while listening grpc host port", logger.Error(err))
		return
	}
	log.Info("Service is running...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Error("error while serving grpc", logger.Error(err))
	}

}
