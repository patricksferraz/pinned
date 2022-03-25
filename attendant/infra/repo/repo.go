package repo

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/attendant/domain/entity"
	"github.com/patricksferraz/attendant/infra/client/kafka"
	"github.com/patricksferraz/attendant/infra/db"
)

type Repository struct {
	Pg *db.PostgreSQL
	Kp *kafka.KafkaProducer
}

func NewRepository(pg *db.PostgreSQL, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Pg: pg,
		Kp: kp,
	}
}

func (r *Repository) CreateAttendant(ctx context.Context, attendant *entity.Attendant) error {
	err := r.Pg.Db.Create(attendant).Error
	return err
}

func (r *Repository) FindAttendant(ctx context.Context, attendantID *string) (*entity.Attendant, error) {
	var e entity.Attendant

	r.Pg.Db.First(&e, "id = ?", *attendantID)

	if e.ID == nil {
		return nil, fmt.Errorf("no attendant found")
	}

	return &e, nil
}

func (r *Repository) SaveAttendant(ctx context.Context, attendant *entity.Attendant) error {
	err := r.Pg.Db.Save(attendant).Error
	return err
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
