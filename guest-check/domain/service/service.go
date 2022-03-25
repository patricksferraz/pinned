package service

import (
	"context"

	"github.com/patricksferraz/guest-check/domain/entity"
	"github.com/patricksferraz/guest-check/domain/repo"
	"github.com/patricksferraz/guest-check/infra/client/kafka/topic"
	"github.com/patricksferraz/guest-check/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateGuest(ctx context.Context, guestID *string) (*string, error) {
	guest, err := entity.NewGuest(guestID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateGuest(ctx, guest); err != nil {
		return nil, err
	}

	return guest.ID, nil
}

func (s *Service) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (s *Service) CreatePlace(ctx context.Context, placeID *string) (*string, error) {
	place, err := entity.NewPlace(placeID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreatePlace(ctx, place); err != nil {
		return nil, err
	}

	return place.ID, nil
}

func (s *Service) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	place, err := s.Repo.FindPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (s *Service) CreateGuestCheck(ctx context.Context, local, guestID, placeID *string) (*string, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	place, err := s.Repo.FindPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	guestCheck, err := entity.NewGuestCheck(local, guest, place)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateGuestCheck(ctx, guestCheck); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(guestCheck)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_GUEST_CHECK), utils.PString(string(eMsg)), guestCheck.ID)
	if err != nil {
		return nil, err
	}

	return guestCheck.ID, nil
}

func (s *Service) FindGuestCheck(ctx context.Context, guestCheckID *string) (*entity.GuestCheck, error) {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return nil, err
	}

	return guestCheck, nil
}

func (s *Service) WaitPaymentGuestCheck(ctx context.Context, guestCheckID *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	err = guestCheck.WaitPayment()
	if err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) CancelGuestCheck(ctx context.Context, guestCheckID, canceledReason *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	err = guestCheck.Cancel(canceledReason)
	if err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) OpenGuestCheck(ctx context.Context, guestCheckID, attendantID *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	attendant, err := s.Repo.FindAttendant(ctx, attendantID)
	if err != nil {
		return err
	}

	if err := guestCheck.Open(attendant); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) PayGuestCheck(ctx context.Context, guestCheckID *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	if err := guestCheck.Pay(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddGuestCheckItem(ctx context.Context, name *string, code, quantity *int, unitPrice *float64, discount *float64, note *string, tag *[]string, guestCheckID *string) (*string, error) {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return nil, err
	}

	guestCheckItem, err := entity.NewGuestCheckItem(name, code, quantity, unitPrice, discount, note, tag, guestCheck)
	if err != nil {
		return nil, err
	}

	err = guestCheck.AddItem(guestCheckItem)
	if err != nil {
		return nil, err
	}

	// TODO: Adds transaction
	err = s.Repo.CreateGuestCheckItem(ctx, guestCheckItem)
	if err != nil {
		return nil, err
	}

	err = s.Repo.SaveGuestCheck(ctx, guestCheck)
	if err != nil {
		return nil, err
	}

	return guestCheckItem.ID, nil
}

func (s *Service) FindGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) (*entity.GuestCheckItem, error) {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return nil, err
	}

	return guestCheckItem, nil
}

func (s *Service) CancelGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID, canceledReason *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Cancel(canceledReason); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) PrepareGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Prepare(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) ReadyGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Ready(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) ForwardGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Forward(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeliverGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Deliver(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateAttendant(ctx context.Context, attendantID *string) (*string, error) {
	attendant, err := entity.NewAttendant(attendantID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateAttendant(ctx, attendant); err != nil {
		return nil, err
	}

	return attendant.ID, nil
}

func (s *Service) FindAttendant(ctx context.Context, attendantID *string) (*entity.Attendant, error) {
	attendant, err := s.Repo.FindAttendant(ctx, attendantID)
	if err != nil {
		return nil, err
	}

	return attendant, nil
}
