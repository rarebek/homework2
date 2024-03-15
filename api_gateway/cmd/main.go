package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"

	"github.com/gomodule/redigo/redis"

	"FreelancerMarketplace/api-gateway/api"
	"FreelancerMarketplace/api-gateway/config"
	"FreelancerMarketplace/api-gateway/pkg/logger"
	"FreelancerMarketplace/api-gateway/services"
	rds "FreelancerMarketplace/api-gateway/storage/redis"
)

// @title App
// @version 0.1
// @description application description
// @securityDefinitions.apikey Token
// @in header
// @name Authorization
func main() {
	//	var redisRepo repo.RedisRepositoryStorage

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	// psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresDatabase,
	// )

	// db, err := gormadapter.NewAdapter("postgres", psqlString, true)
	// if err != nil {
	// 	log.Error("new adapter error", logger.Error(err))
	// 	return
	// }

	// casbinEnforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, db)
	// if err != nil {
	// 	log.Error("new enforcer error", logger.Error(err))
	// 	return
	// }

	// err = casbinEnforcer.LoadPolicy()
	// if err != nil {
	// 	log.Error("new load policy error", logger.Error(err))
	// 	return
	// }
	pool := redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	redisRepo := rds.NewRedisRepo(&pool)

	// casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch", util.KeyMatch)
	// casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch3", util.KeyMatch3)

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		Casbin:         &casbin.Enforcer{},
		ServiceManager: serviceManager,
		RedisRepo:      redisRepo,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
