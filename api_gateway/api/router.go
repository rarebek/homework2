package api

import (
	casbinN "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "FreelancerMarketplace/api-gateway/api/docs"
	v1 "FreelancerMarketplace/api-gateway/api/handlers/v1"
	"FreelancerMarketplace/api-gateway/config"
	"FreelancerMarketplace/api-gateway/pkg/logger"
	"FreelancerMarketplace/api-gateway/services"
	"FreelancerMarketplace/api-gateway/storage/repo"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	RedisRepo      repo.RedisRepositoryStorage
	Casbin         *casbinN.Enforcer
}

// New ...
func New(option Option) *gin.Engine {
	enforcer, _ := casbinN.NewEnforcer(option.Conf.CasbinConfigPath, option.Conf.AuthCSVPath)
	_ = enforcer.LoadPolicy()
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.RedisRepo,
		Casbin:         option.Casbin,
	})

	api := router.Group("/v1")

	// api.Use(casbin.NewAuth(enforcer, option.Conf))

	api.POST("/company/register", handlerV1.RegisterCompany)
	api.POST("/company/verify/:email/:code", handlerV1.VerifyCompany)
	api.POST("/company/login/:email/:password", handlerV1.LogInCompany)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
