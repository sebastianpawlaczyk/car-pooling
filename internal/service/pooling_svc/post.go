package pooling_svc

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/car-pooling/internal/model"
)

var CanNotLocateJourneyError = errors.New("can not local journey")

func (svc *poolingService) PostJourney(ctx context.Context, journey *model.Journey) error {
	availableCar, err := svc.repository.GetCarByAvailableSeats(ctx, journey.People)
	if err != nil {
		return err
	}

	if availableCar != nil {
		remainingSeats := availableCar.Seats - journey.People
		availableCar.Seats = remainingSeats
		if err := svc.repository.UpdateCar(ctx, availableCar); err != nil {
			return err
		}

		journey.CarID = strconv.Itoa(availableCar.ID)
	}

	journey.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	err = svc.repository.InsertJourney(ctx, journey)
	if err != nil {
		return err
	}

	return nil
}

func (svc *poolingService) PostDropoff(ctx context.Context, journeyID int) (*model.Journey, error) {
	journey, err := svc.repository.GetJourneyByID(ctx, journeyID)
	if err != nil {
		return nil, err
	}

	if journey != nil {
		if err := svc.repository.DeleteJourneyByID(ctx, journeyID); err != nil {
			return nil, err
		}

		if journey.CarID != "" {
			carID, err := strconv.Atoi(journey.CarID)
			if err != nil {
				return nil, err
			}
			car, err := svc.repository.GetCarByID(ctx, carID)
			if err != nil {
				return nil, err
			}
			if car == nil {
				return nil, fmt.Errorf("can not find assigned car")
			}

			updatedSeats := car.Seats + journey.People
			car.Seats = updatedSeats

			// check if some waiting group matches to dropped car
			car, err = svc.checkQueue(ctx, car)
			if err != nil {
				return nil, err
			}

			if err := svc.repository.UpdateCar(ctx, car); err != nil {
				return nil, err
			}
		}
	}

	return journey, nil
}

func (svc *poolingService) PostLocate(ctx context.Context, journeyID int) (*model.Car, error) {
	journey, err := svc.repository.GetJourneyByID(ctx, journeyID)
	if err != nil {
		return nil, err
	}

	if journey == nil {
		return nil, CanNotLocateJourneyError
	}

	if journey.CarID != "" {
		carID, err := strconv.Atoi(journey.CarID)
		if err != nil {
			return nil, err
		}
		car, err := svc.repository.GetCarByID(ctx, carID)
		if err != nil {
			return nil, err
		}

		if car == nil {
			return nil, fmt.Errorf("can not find assigned car")
		}

		assignedJourneys, err := svc.repository.GetJourneysByCarID(ctx, journey.CarID)
		if err != nil {
			return nil, fmt.Errorf("error during finding journeys")
		}

		seats := car.Seats
		for _, j := range assignedJourneys {
			seats += j.People
		}
		response := &model.Car{
			ID:    car.ID,
			Seats: seats,
		}

		return response, nil
	}

	return nil, nil
}
