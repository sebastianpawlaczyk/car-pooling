package pooling_api

import (
	"errors"

	"github.com/car-pooling/internal/model"
)

var TooManySeatsInCarError = errors.New("too many seats in a car")
var TooBigGroupError = errors.New("group is too big")

func validCars(cars []*model.Car) error {
	for _, car := range cars {
		if car.Seats > 6 {
			return TooManySeatsInCarError
		}
	}

	return nil
}

func validJourney(journey *model.Journey) error {
	if journey.People > 6 {
		return TooBigGroupError
	}

	return nil
}
