package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Consumer       *ckafka.Consumer
	ConsumerTopics []string
}

func NewKafkaConsumer(servers, groupId string, consumerTopics []string) (*KafkaConsumer, error) {
	c, err := ckafka.NewConsumer(
		&ckafka.ConfigMap{
			"bootstrap.servers": servers,
			"group.id":          groupId,
			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		Consumer:       c,
		ConsumerTopics: consumerTopics,
	}, nil
}
