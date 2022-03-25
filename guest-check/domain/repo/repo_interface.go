package repo

import (
	"context"

	"github.com/patricksferraz/guest-check/domain/entity"
)

type RepoInterface interface {
	CreateGuest(ctx context.Context, guest *entity.Guest) error
	FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error)
	SaveGuest(ctx context.Context, guest *entity.Guest) error

	CreatePlace(ctx context.Context, place *entity.Place) error
	FindPlace(ctx context.Context, placeID *string) (*entity.Place, error)
	SavePlace(ctx context.Context, place *entity.Place) error

	CreateGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error
	FindGuestCheck(ctx context.Context, guestCheckID *string) (*entity.GuestCheck, error)
	SaveGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error

	CreateGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error
	FindGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) (*entity.GuestCheckItem, error)
	SaveGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error

	CreateAttendant(ctx context.Context, attendant *entity.Attendant) error
	FindAttendant(ctx context.Context, attendantID *string) (*entity.Attendant, error)
	SaveAttendant(ctx context.Context, attendant *entity.Attendant) error

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
