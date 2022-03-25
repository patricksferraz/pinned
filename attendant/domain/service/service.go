package service

import (
	"context"

	"github.com/patricksferraz/attendant/domain/entity"
	"github.com/patricksferraz/attendant/domain/repo"
	"github.com/patricksferraz/attendant/infra/client/kafka/topic"
	"github.com/patricksferraz/attendant/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateAttendant(ctx context.Context, name *string) (*string, error) {
	attendant, err := entity.NewAttendant(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateAttendant(ctx, attendant); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(attendant)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_ATTENDANT), utils.PString(string(eMsg)), attendant.ID)
	if err != nil {
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
