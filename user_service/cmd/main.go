package main

import (
	"FreelancerMarketplace/user-service/config"
	pb "FreelancerMarketplace/user-service/genproto/user_service"
	"FreelancerMarketplace/user-service/pkg/db"
	"FreelancerMarketplace/user-service/pkg/logger"
	consumer2 "FreelancerMarketplace/user-service/queue/kafka/consumer"
	"FreelancerMarketplace/user-service/service"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	log = logger.WithFields(log, logger.String("file", "log.txt"))

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres error", logger.Error(err))
	}

	authorizationService := service.NewUserService(connDB, log)

	consumer, err := consumer2.NewKafkaConsumer([]string{"kafka:9092"}, "test-topic", "nodirbek", authorizationService)
	if err != nil {
		fmt.Println(err)
	}
	defer consumer.Close()

	go func() {
		consumer.ConsumeMessages(consumer2.ConsumeHandler)
	}()

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterCompanyServiceServer(s, authorizationService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

func consumerHandler(message []byte) {
	fmt.Println(string(message))
}
