package services

import (
	"fmt"

	"FreelancerMarketplace/api-gateway/config"
	pb "FreelancerMarketplace/api-gateway/genproto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	CompanyService() pb.CompanyServiceClient
}

type serviceManager struct {
	companyService pb.CompanyServiceClient
}

func (s *serviceManager) CompanyService() pb.CompanyServiceClient {
	return s.companyService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		companyService: pb.NewCompanyServiceClient(connUser),
	}

	return serviceManager, nil
}
