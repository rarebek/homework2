package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "FreelancerMarketplace/user-service/genproto/user_service"
	"FreelancerMarketplace/user-service/service"

	"github.com/k0kubun/pp"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	ConsumeMessages(handler func(message []byte, userService *service.UserService)) error
	Close() error
}

type Consumer struct {
	reader      *kafka.Reader
	userService *service.UserService
}

func NewKafkaConsumer(brokers []string, topic string, groupID string, userService *service.UserService) (KafkaConsumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return &Consumer{reader: reader, userService: userService}, nil
}

func (c *Consumer) ConsumeMessages(handler func(message []byte, userService *service.UserService)) error {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		fmt.Println(err, "++++")
		if err != nil {
			return err
		}
		handler(msg.Value, c.userService)
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func ConsumeHandler(message []byte, service *service.UserService) {
	fmt.Println(string(message))
	var user pb.CompanyRegisterRequest
	if err := json.Unmarshal(message, &user); err != nil {
		log.Fatal("cannot unmarshal json")
		return
	}

	//respUser, err := service.CompanyRegister(context.Background(), &user)
	//if err != nil {
	//	log.Fatal("cannot create user via kafka")
	//	return
	//}
	pp.Println(user)
}
