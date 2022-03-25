package repo

import (
	"context"

	"github.com/patricksferraz/menu/domain/entity"
)

type RepoInterface interface {
	CreateMenu(ctx context.Context, menu *entity.Menu) error
	FindMenu(ctx context.Context, menuID *string) (*entity.Menu, error)
	SaveMenu(ctx context.Context, menu *entity.Menu) error

	CreateItem(ctx context.Context, item *entity.Item) error
	FindItem(ctx context.Context, menuID, itemID *string) (*entity.Item, error)
	SaveItem(ctx context.Context, item *entity.Item) error

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
