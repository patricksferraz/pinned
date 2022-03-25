package repo

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/menu/domain/entity"
	"github.com/patricksferraz/menu/infra/client/kafka"
	"github.com/patricksferraz/menu/infra/db"
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

func (r *Repository) CreateMenu(ctx context.Context, menu *entity.Menu) error {
	err := r.Pg.Db.Create(menu).Error
	return err
}

func (r *Repository) FindMenu(ctx context.Context, menuID *string) (*entity.Menu, error) {
	var e entity.Menu

	r.Pg.Db.First(&e, "id = ?", *menuID)

	if e.ID == nil {
		return nil, fmt.Errorf("no menu found")
	}

	return &e, nil
}

func (r *Repository) SaveMenu(ctx context.Context, menu *entity.Menu) error {
	err := r.Pg.Db.Save(menu).Error
	return err
}

func (r *Repository) CreateItem(ctx context.Context, item *entity.Item) error {
	err := r.Pg.Db.Create(item).Error
	return err
}

func (r *Repository) FindItem(ctx context.Context, menuID, itemID *string) (*entity.Item, error) {
	var e entity.Item

	r.Pg.Db.First(&e, "id = ? AND menu_id = ?", *itemID, *menuID)

	if e.ID == nil {
		return nil, fmt.Errorf("no item found")
	}

	return &e, nil
}

func (r *Repository) SaveItem(ctx context.Context, item *entity.Item) error {
	err := r.Pg.Db.Save(item).Error
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
