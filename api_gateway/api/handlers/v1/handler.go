package v1

import (
	"FreelancerMarketplace/api-gateway/config"
	"FreelancerMarketplace/api-gateway/pkg/logger"
	"FreelancerMarketplace/api-gateway/services"
	"FreelancerMarketplace/api-gateway/storage/repo"
	"github.com/casbin/casbin/v2"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redisStorage   repo.RedisRepositoryStorage
	casbin         *casbin.Enforcer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.RedisRepositoryStorage
	Casbin         *casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redisStorage:   c.Redis,
		casbin:         c.Casbin,
	}
}
