package sqlc

import (
	pb "carpet/genproto/carpet_service"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Querier interface {
	InsertCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error)
	DeleteCompany(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	SelectCompany(ctx context.Context, req *pb.PrimaryKey) (*pb.Company, error)
	UpdateCompany(ctx context.Context, req *pb.Company) (*pb.Company, error)
	InsertAddress(ctx context.Context, req *pb.AddressRequest) (*pb.Address, error)
	UpdateAddress(ctx context.Context, req *pb.Address) (*pb.Address, error)
	DeleteAddress(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	SelectAddress(ctx context.Context, req *pb.PrimaryKey) (*pb.Address, error)
	InsertOrder(ctx context.Context, req *pb.OrderRequest) (*pb.Order, error)
	UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error)
	//UpdateOrderWithUser(ctx context.Context, req *pb.Order) (*pb.Order, error)
	DeleteOrder(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	SelectOrders(ctx context.Context, req *pb.GetListRequest) (*pb.OrdersResponse, error)
	SelectOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.Order, error)

	InsertService(ctx context.Context, req *pb.ServiceRequest) (*pb.Service, error)
	UpdateService(ctx context.Context, req *pb.Service) (*pb.Service, error)
	DeleteService(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error)
	SelectService(ctx context.Context, req *pb.PrimaryKey) (*pb.Service, error)
	SelectServices(ctx context.Context, req *pb.GetListRequest) (*pb.ServicesResponse, error)
}

var _ Querier = (*Queries)(nil)
