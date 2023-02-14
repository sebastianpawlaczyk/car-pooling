package pooling_svc

import (
	"context"

	"github.com/car-pooling/internal/model"
)

func (svc *poolingService) checkQueue(ctx context.Context, car *model.Car) (*model.Car, error) {
	for {
		next, err := svc.repository.UpdateNextPossibleJourney(ctx, car)
		if err != nil {
			return nil, err
		}
		if next == nil {
			break
		}
		car.Seats -= next.People
	}

	return car, nil
}
