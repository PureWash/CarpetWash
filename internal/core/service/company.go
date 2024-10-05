package service

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/pkg/logger"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type Company struct {
	storage sqlc.Querier
	log     logger.ILogger
	pb.UnimplementedCompanyServiceServer
}

func NewCompany(storage sqlc.Querier, log logger.ILogger) *Company {
	return &Company{
		storage: storage,
		log:     log,
	}
}

func (s *Company) CreateCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error) {
	s.log.Info("Insert company successfully")
	res, err := s.storage.InsertCompany(ctx, req)
	if err != nil {
		s.log.Error("Insert company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}

func (s *Company) UpdateCompany(ctx context.Context, req *pb.Company) (*pb.Company, error) {
	s.log.Info("Update company successfully")
	res, err := s.storage.UpdateCompany(ctx, req)
	if err != nil {
		s.log.Error("Update company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}

func (s *Company) DeleteCompany(ctx context.Context, req *pb.PrimaryKey) (*empty.Empty, error) {
	s.log.Info("Delete company successfully")
	res, err := s.storage.DeleteCompany(ctx, req)
	if err != nil {
		s.log.Error("Delete company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}

func (s *Company) GetCompany(ctx context.Context, req *pb.PrimaryKey) (*pb.Company, error) {
	s.log.Info("Select company successfully")
	res, err := s.storage.SelectCompany(ctx, req)
	if err != nil {
		s.log.Error("Select company error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success company service")
	return res, nil
}
