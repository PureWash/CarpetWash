package service

import (
	pb "carpet/genproto/carpet_service"
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/pkg/logger"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type Order struct {
	storage sqlc.Querier
	log     logger.ILogger
	pb.UnimplementedOrderServiceServer
}

func NewOrder(storage sqlc.Querier, log logger.ILogger) *Order {
	return &Order{
		storage: storage,
		log:     log,
	}
}

func (s *Order) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.Order, error) {
	s.log.Info("Insert Order successfully")
	res, err := s.storage.InsertOrder(ctx, req)
	if err != nil {
		s.log.Error("Insert Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	s.log.Info("Update Order successfully")
	res, err := s.storage.UpdateOrder(ctx, req)
	if err != nil {
		s.log.Error("Update Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) UpdateOrderWithUser(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	s.log.Info("Update Order with user successfully")
	res, err := s.storage.UpdateOrderWithUser(ctx, req)
	if err != nil {
		s.log.Error("Update Order with user error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) DeleteOrder(ctx context.Context, req *pb.PrimaryKey) (*empty.Empty, error) {
	s.log.Info("Delete Order successfully")
	res, err := s.storage.DeleteOrder(ctx, req)
	if err != nil {
		s.log.Error("Delete Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) GetOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.Order, error) {
	s.log.Info("Select Order successfully")
	res, err := s.storage.SelectOrder(ctx, req)
	if err != nil {
		s.log.Error("Select Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) GetAllOrder(ctx context.Context, req *pb.GetListRequest) (*pb.OrdersResponse, error) {
	s.log.Info("Select Orders successfully")
	res, err := s.storage.SelectOrders(ctx, req)
	if err != nil {
		s.log.Error("Select Orders error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}
