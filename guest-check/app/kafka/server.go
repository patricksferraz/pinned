package kafka

import (
	"fmt"

	"github.com/patricksferraz/guest-check/domain/service"
	"github.com/patricksferraz/guest-check/infra/client/kafka"
	"github.com/patricksferraz/guest-check/infra/db"
	"github.com/patricksferraz/guest-check/infra/repo"
)

func StartKafkaServer(pg *db.PostgreSQL, kc *kafka.KafkaConsumer, kp *kafka.KafkaProducer) {
	repository := repo.NewRepository(pg, kp)
	service := service.NewService(repository)

	fmt.Println("kafka pocessor has been started")
	processor := NewKafkaProcessor(service, kc)
	processor.Consume()
}
