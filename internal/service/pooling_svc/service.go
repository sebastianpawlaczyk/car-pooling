package pooling_svc

import (
	"context"

	"github.com/car-pooling/internal/model"
	"github.com/car-pooling/internal/repository/pooling_repo"
)

type PoolingInterface interface {
	PutCars(ctx context.Context, cars []*model.Car) error
	PostJourney(ctx context.Context, journey *model.Journey) error
	PostDropoff(ctx context.Context, journeyID int) (*model.Journey, error)
	PostLocate(ctx context.Context, journeyID int) (*model.Car, error)
}

type poolingService struct {
	repository pooling_repo.PoolingRepository
}

func NewPoolingService(repository pooling_repo.PoolingRepository) PoolingInterface {
	return &poolingService{
		repository: repository,
	}
}
