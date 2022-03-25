package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProducer(servers string, deliveryChan chan ckafka.Event) (*KafkaProducer, error) {
	p, err := ckafka.NewProducer(
		&ckafka.ConfigMap{
			"bootstrap.servers": servers,
		},
	)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Producer:     p,
		DeliveryChan: deliveryChan,
	}, nil
}

// TODO: Add event log
func (k *KafkaProducer) DeliveryReport() {
	for e := range k.DeliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				// TODO: add attempts
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
