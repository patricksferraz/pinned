package kafka

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/guest-check/app/kafka/event"
	"github.com/patricksferraz/guest-check/domain/service"
	"github.com/patricksferraz/guest-check/infra/client/kafka"
	"github.com/patricksferraz/guest-check/infra/client/kafka/topic"
)

type KafkaProcessor struct {
	Service *service.Service
	Kc      *kafka.KafkaConsumer
}

func NewKafkaProcessor(service *service.Service, kafkaConsumer *kafka.KafkaConsumer) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		Kc:      kafkaConsumer,
	}
}

func (p *KafkaProcessor) Consume() {
	p.Kc.Consumer.SubscribeTopics(p.Kc.ConsumerTopics, nil)
	for {
		msg, err := p.Kc.Consumer.ReadMessage(-1)
		if err == nil {
			// fmt.Println(string(msg.Value))
			err := p.processMessage(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) error {
	switch _topic := *msg.TopicPartition.Topic; _topic {
	// GUEST_CHECK
	case topic.OPEN_GUEST_CHECK:
		err := p.openGuestCheck(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("open guest check, error %s", err)
		}
	// GUEST
	case topic.NEW_GUEST:
		err := p.createGuest(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create guest, error %s", err)
		}
	// PLACE
	case topic.NEW_PLACE:
		err := p.createPlace(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create place, error %s", err)
		}
	// ATTENDANT
	case topic.NEW_ATTENDANT:
		err := p.createAttendant(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create attendant, error %s", err)
		}
	default:
		return fmt.Errorf("not a valid topic %s", string(msg.Value))
	}

	return nil
}

func (p *KafkaProcessor) retry(msg *ckafka.Message) error {
	err := p.Kc.Consumer.Seek(ckafka.TopicPartition{
		Topic:     msg.TopicPartition.Topic,
		Partition: msg.TopicPartition.Partition,
		Offset:    msg.TopicPartition.Offset,
	}, -1)

	return err
}

func (p *KafkaProcessor) openGuestCheck(msg *ckafka.Message) error {
	e := &event.OpenGuestCheck{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.OpenGuestCheck(context.TODO(), e.Msg.GuestCheckID, e.Msg.AttendantID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createGuest(msg *ckafka.Message) error {
	e := &event.Guest{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateGuest(context.TODO(), e.Msg.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createPlace(msg *ckafka.Message) error {
	e := &event.Place{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreatePlace(context.TODO(), e.Msg.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createAttendant(msg *ckafka.Message) error {
	e := &event.Attendant{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateAttendant(context.TODO(), e.Msg.ID)
	if err != nil {
		return err
	}

	return nil
}
