package service

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/pkg/logger"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) InsertOrder(ctx context.Context, req *pb.OrderRequest) (*pb.Order, error) {
	s.log.Info("Insert Order successfully")
	res, err := s.storage.InsertOrder(ctx, req)
	if err != nil {
		s.log.Error("Insert Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	s.log.Info("Update Order successfully")
	res, err := s.storage.UpdateOrder(ctx, req)
	if err != nil {
		s.log.Error("Update Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Service) UpdateOrderWithUser(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	s.log.Info("Update Order with user successfully")
	res, err := s.storage.UpdateOrderWithUser(ctx, req)
	if err != nil {
		s.log.Error("Update Order with user error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("Delete Order successfully")
	res, err := s.storage.DeleteOrder(ctx, req)
	if err != nil {
		s.log.Error("Delete Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Service) SelectOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.Order, error) {
	s.log.Info("Select Order successfully")
	res, err := s.storage.SelectOrder(ctx, req)
	if err != nil {
		s.log.Error("Select Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Service) SelectOrders(ctx context.Context, req *pb.GetListRequest) (*pb.OrdersResponse, error) {
	s.log.Info("Select Orders successfully")
	res, err := s.storage.SelectOrders(ctx, req)
	if err != nil {
		s.log.Error("Select Orders error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}
