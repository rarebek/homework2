package repo

import (
	pb "FreelancerMarketplace/user-service/genproto/user_service"
)

// UserStorageI ...
type UserStorageI interface {
	CompanyRegister(request *pb.CompanyRegisterRequest) (*pb.CompanyRegisterResponse, error)
	CompanyLogIn(request *pb.CompanyLogInRequest) (*pb.CompanyLogInResponse, error)
	CheckUniqueness(request *pb.CheckUniquenessRequest) (*pb.CheckUniquenessResponse, error)
}
