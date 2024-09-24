package grpcConn

import (
	carPet "carpet/genproto/pure_wash"
	"carpet/internal/core/repository/psql/sqlc"

	"google.golang.org/grpc"
)

func ConnGRPC(storega *sqlc.Querier) {
	grpcServer := grpc.NewServer()

	carPet.RegisterAddressesServer(grpcServer, )
}
