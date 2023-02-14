package pooling_svc

import (
	"context"

	"github.com/car-pooling/internal/model"
)

func (svc *poolingService) PutCars(ctx context.Context, cars []*model.Car) error {
	err := svc.repository.DeleteJourneys(ctx)
	if err != nil {
		return err
	}

	err = svc.repository.DeleteCars(ctx)
	if err != nil {
		return err
	}

	return svc.repository.InsertCars(ctx, cars)
}
