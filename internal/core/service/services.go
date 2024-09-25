package service

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/pkg/logger"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	storage sqlc.Service
	log     logger.ILogger
	pb.UnimplementedServiceServiceServer
}

func NewService(storage sqlc.Service, log logger.ILogger) *Service {
	return &Service{
		storage: storage,
		log:     log,
	}
}

func (s *Service) InsertService(ctx context.Context, req *pb.ServiceRequest) (*pb.Service, error) {
	s.log.Info("Insert service successfully")
	res, err := s.storage.InsertService(ctx, req)
	if err != nil {
		s.log.Error("Insert service error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success insert")
	return res, nil
}

func (s *Service) UpdateService(ctx context.Context, req *pb.Service) (*pb.Service, error) {
	s.log.Info("Update service successfully")
	res, err := s.storage.UpdateService(ctx, req)
	if err != nil {
		s.log.Error("Update service error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success update")
	return res, nil
}

func (s *Service) DeleteService(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("Delete service successfully")
	res, err := s.storage.DeleteService(ctx, req)
	if err != nil {
		s.log.Error("Delete service error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success delete")
	return res, nil
}

func (s *Service) SelectService(ctx context.Context, req *pb.PrimaryKey) (*pb.Service, error) {
	s.log.Info("Select service successfully")
	res, err := s.storage.SelectService(ctx, req)
	if err != nil {
		s.log.Error("Select service error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success select service")
	return res, nil
}

func (s *Service) SelectServices(ctx context.Context, req *pb.GetListRequest) (*pb.ServicesResponse, error) {
	s.log.Info("Select services successfully")
	res, err := s.storage.SelectServices(ctx, req)
	if err != nil {
		s.log.Error("Select service error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success select services")
	return res, nil
}
