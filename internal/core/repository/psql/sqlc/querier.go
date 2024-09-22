package sqlc

import (
	pb "carpet/genproto/pure_wash"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Querier interface {
	InsertCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error)
	InsertService(ctx context.Context, req *pb.ServiceRequest) (*pb.Service, error)
	InsertAddress(ctx context.Context, req *pb.AddressRequest) (*pb.Address, error)
	InsertOrder(ctx context.Context, req *pb.OrderRequest) (*pb.Order, error)
	UpdateCompany(ctx context.Context, req *pb.Company) (*pb.Company, error)
	UpdateService(ctx context.Context, req *pb.Service) (*pb.Service, error)
	UpdateAddress(ctx context.Context, req *pb.Address) (*pb.Address, error)
	UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error)
	DeleteCompany(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	DeleteService(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	DeleteAddress(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	DeleteOrder(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	SelectServices(ctx context.Context, req *pb.GetListRequest) (*pb.ServicesResponse, error)
	SelectOrders(ctx context.Context, req *pb.GetListRequest) (*pb.OrdersResponse, error)
	SelectCompany(ctx context.Context, req *pb.PrimaryKey) (*pb.Company, error)
	SelectService(ctx context.Context, req *pb.PrimaryKey) (*pb.Service, error)	
	SelectAddress(ctx context.Context, req *pb.PrimaryKey) (*pb.Address, error)
	SelectOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.Order, error)
}

var _, Querier = (*Queries)(nil)
