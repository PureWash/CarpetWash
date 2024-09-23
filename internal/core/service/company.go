package service

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/pkg/logger"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) InsertCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error) {
	s.log.Info("Insert company successfully")
	res, err := s.storage.InsertCompany(ctx, req)
	if err != nil {
		s.log.Error("Insert company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}

func (s *Service) UpdateCompany(ctx context.Context, req *pb.Company) (*pb.Company, error) {
	s.log.Info("Update company successfully")
	res, err := s.storage.UpdateCompany(ctx, req)
	if err != nil {
		s.log.Error("Update company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}

func (s *Service) DeleteCompany(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("Delete company successfully")
	res, err := s.storage.DeleteCompany(ctx, req)
	if err != nil {
		s.log.Error("Delete company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}

func (s *Service) SelectCompany(ctx context.Context, req *pb.PrimaryKey) (*pb.Company, error) {
	s.log.Info("Select company successfully")
	res, err := s.storage.SelectCompany(ctx, req)
	if err != nil {
		s.log.Error("Select company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}
