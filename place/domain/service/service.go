package service

import (
	"context"

	"github.com/patricksferraz/place/domain/entity"
	"github.com/patricksferraz/place/domain/repo"
	"github.com/patricksferraz/place/infra/client/kafka/topic"
	"github.com/patricksferraz/place/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreatePlace(ctx context.Context, name *string) (*string, error) {
	place, err := entity.NewPlace(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreatePlace(ctx, place); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(place)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_PLACE), utils.PString(string(eMsg)), place.ID)
	if err != nil {
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
