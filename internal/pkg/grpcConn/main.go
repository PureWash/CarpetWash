package grpcConn

import (
	carPet "carpet/genproto/pure_wash"
	"carpet/internal/core/repository/psql/sqlc"
	service2 "carpet/internal/core/service"
	"carpet/internal/pkg/logger"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func ConnGRPC(storage sqlc.Querier, log logger.ILogger) *grpc.Server {
	grpcServer := grpc.NewServer()

	carPet.RegisterAddressesServer(grpcServer, service2.NewAddress(storage, log))
	carPet.RegisterCompanyServiceServer(grpcServer, service2.NewCompany(storage, log))
	carPet.RegisterOrderServiceServer(grpcServer, service2.NewOrder(storage, log))
	carPet.RegisterServiceServiceServer(grpcServer, service2.NewService(storage, log))
	reflection.Register(grpcServer)

	return grpcServer
}
