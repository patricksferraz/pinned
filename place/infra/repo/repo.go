package repo

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/place/domain/entity"
	"github.com/patricksferraz/place/infra/client/kafka"
	"github.com/patricksferraz/place/infra/db"
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

func (r *Repository) CreatePlace(ctx context.Context, place *entity.Place) error {
	err := r.Pg.Db.Create(place).Error
	return err
}

func (r *Repository) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	var e entity.Place

	r.Pg.Db.First(&e, "id = ?", *placeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no place found")
	}

	return &e, nil
}

func (r *Repository) SavePlace(ctx context.Context, place *entity.Place) error {
	err := r.Pg.Db.Save(place).Error
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
