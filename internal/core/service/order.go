package service

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/models"
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

func (s *Order) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	s.log.Info("Insert Order successfully")

	resp1, err := s.storage.CreateClient(ctx, models.CreateClientReq{
		FullName:    req.Client.FullName,
		PhoneNumber: req.Client.PhoneNumber,
		Latitude:    req.Client.Latitude,
		Longitude:   req.Client.Longitude,
	})

	if err != nil {
		return nil, err
	}

	resp, err := s.storage.InsertOrder(ctx, models.CreateOrderReq{
		ClientID:   resp1.ID,
		Area:       req.Area,
		TotalPrice: req.TotalPrice,
		ServiceId:  req.ServiceId,
	})
	if err != nil {
		s.log.Error("Insert Order error", logger.Error(err))
		return nil, err
	}

	s.log.Info("Success Order service")
	return &pb.CreateOrderResp{
		FullName:    resp1.FullName,
		PhoneNumber: resp1.PhoneNumber,
		Area:        resp.Area,
		TotalPrice:  resp.TotalPrice,
		CreatedAt:   resp.CreatedAt,
	}, nil
}

func (s *Order) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	s.log.Info("Update Order successfully")

	res, err := s.storage.UpdateOrder(ctx, models.UpdateOrderReq{
		ID:         req.GetId(),
		Area:       req.GetArea(),
		TotalPrice: req.GetTotalPrice(),
		Status:     req.GetStatus(),
	})
	if err != nil {
		s.log.Error("Update Order error", logger.Error(err))
		return nil, err
	}

	err = s.storage.UpdateClient(ctx, models.UpdateClientReq{
		ID:          res.ClientID,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	s.log.Info("Success Order service")
	return &pb.UpdateOrderResp{
		Id:         res.ID,
		Area:       res.Area,
		TotalPrice: res.TotalPrice,
		UpdatedAt:  res.UpdatedAt,
	}, nil
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

func (s *Order) GetOrder(ctx context.Context, req *pb.PrimaryKey) (*pb.GetOrderResp, error) {
	s.log.Info("Select Order successfully")
	res, err := s.storage.SelectOrder(ctx, req)
	if err != nil {
		s.log.Error("Select Order error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) GetAllOrder(ctx context.Context, req *pb.GetListRequest) (*pb.GetOrdersResp, error) {
	s.log.Info("Select Orders successfully")
	res, err := s.storage.SelectOrders(ctx, req)
	if err != nil {
		s.log.Error("Select Orders error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order service")
	return res, nil
}

func (s *Order) GetAllOrderForCurier(ctx context.Context, req *pb.GetAllOrdersReq) (*pb.GetOrdersResp, error) {
	s.log.Info("Get Orders filter successfully")
	resp, err := s.storage.GetAllOrders(ctx, req)
	if err != nil {
		s.log.Error("Select Order filters error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success Order filter service")
	return resp, err
}

func (s *Order) UpdateOrderStatus(ctx context.Context, in *pb.StatusOrderReq) (*pb.PrimaryKey, error) {
	res, err := s.storage.UpdateOrderStatus(ctx, in)
	if err != nil {
		s.log.Error("Select service error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success select services")
	return res, nil
}
