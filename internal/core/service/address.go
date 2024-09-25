package service

import (
	pb "carpet/genproto/pure_wash"
	"carpet/internal/core/repository/psql/sqlc"
	"carpet/internal/pkg/logger"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Address struct {
	storage sqlc.Address
	log     logger.ILogger
	pb.UnimplementedAddressesServer
}

func NewAddress(storage sqlc.Address, log logger.ILogger) *Address {
	return &Address{
		storage: storage,
		log:     log,
	}
}

func (s *Address) InsertAddress(ctx context.Context, req *pb.AddressRequest) (*pb.Address, error) {
	s.log.Info("Insert address successfully")
	res, err := s.storage.InsertAddress(ctx, req)
	if err != nil {
		s.log.Error("Insert address error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success address service")
	return res, nil
}

func (s *Address) UpdateAddress(ctx context.Context, req *pb.Address) (*pb.Address, error) {
	s.log.Info("Update address successfully")
	res, err := s.storage.UpdateAddress(ctx, req)
	if err != nil {
		s.log.Error("Update address error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success address service")
	return res, nil
}

func (s *Address) DeleteAddress(ctx context.Context, req *pb.PrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("Delete address successfully")
	res, err := s.storage.DeleteAddress(ctx, req)
	if err != nil {
		s.log.Error("Delete address error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success address service")
	return res, nil
}

func (s *Address) SelectAddress(ctx context.Context, req *pb.PrimaryKey) (*pb.Address, error) {
	s.log.Info("Select address successfully")
	res, err := s.storage.SelectAddress(ctx, req)
	if err != nil {
		s.log.Error("Select address error", logger.Error(err))
		return nil, err
	}
	s.log.Info("Success address service")
	return res, nil
}
