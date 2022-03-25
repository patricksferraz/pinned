package repo

import (
	"context"

	"github.com/patricksferraz/guest/domain/entity"
)

type RepoInterface interface {
	CreateGuest(ctx context.Context, guest *entity.Guest) error
	FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error)
	SaveGuest(ctx context.Context, guest *entity.Guest) error

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
