package service

import (
	pb "FreelancerMarketplace/user-service/genproto/user_service"
	l "FreelancerMarketplace/user-service/pkg/logger"
	"FreelancerMarketplace/user-service/storage"
	"context"
	"database/sql"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedCompanyServiceServer
}

// NewUserService ...
func NewUserService(db *sql.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db, log),
		logger:  log,
	}
}

func (s *UserService) CompanyRegister(ctx context.Context, req *pb.CompanyRegisterRequest) (*pb.CompanyRegisterResponse, error) {
	return s.storage.User().CompanyRegister(req)
}

func (s *UserService) CompanyLogIn(ctx context.Context, req *pb.CompanyLogInRequest) (*pb.CompanyLogInResponse, error) {
	return s.storage.User().CompanyLogIn(req)
}

func (s *UserService) CheckUniqueness(ctx context.Context, req *pb.CheckUniquenessRequest) (*pb.CheckUniquenessResponse, error) {
	return s.storage.User().CheckUniqueness(req)
}
