package pooling_svc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/car-pooling/internal/model"
	mocks "github.com/car-pooling/mocks/mocks_repo"
)

func TestPutCars(t *testing.T) {
	testCases := []struct {
		name              string
		cars              []*model.Car
		deleteJourneysErr error
		deleteCarsErr     error
		insertCarsErr     error
		responseError     error
	}{
		{
			name: "ok",
			cars: []*model.Car{
				{
					ID:    2,
					Seats: 4,
				},
			},
			deleteJourneysErr: nil,
			deleteCarsErr:     nil,
			insertCarsErr:     nil,
			responseError:     nil,
		},
		{
			name: "delete journeys error",
			cars: []*model.Car{
				{
					ID:    2,
					Seats: 4,
				},
			},
			deleteJourneysErr: fmt.Errorf("deleteJourneysErr"),
			deleteCarsErr:     nil,
			insertCarsErr:     nil,
			responseError:     fmt.Errorf("deleteJourneysErr"),
		},
		{
			name: "insert cars error",
			cars: []*model.Car{
				{
					ID:    2,
					Seats: 4,
				},
			},
			deleteJourneysErr: nil,
			deleteCarsErr:     fmt.Errorf("deleteCarsErr"),
			insertCarsErr:     nil,
			responseError:     fmt.Errorf("deleteCarsErr"),
		},
		{
			name: "delete journeys error",
			cars: []*model.Car{
				{
					ID:    2,
					Seats: 4,
				},
			},
			deleteJourneysErr: nil,
			deleteCarsErr:     nil,
			insertCarsErr:     fmt.Errorf("insertCarsErr"),
			responseError:     fmt.Errorf("insertCarsErr"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mockRepo := &mocks.PoolingRepository{}
			svc := NewPoolingService(mockRepo)

			mockRepo.On("DeleteJourneys", mock.Anything).Return(testCase.deleteJourneysErr)
			mockRepo.On("DeleteCars", mock.Anything).Return(testCase.deleteCarsErr)
			mockRepo.On("InsertCars", mock.Anything, mock.Anything).Return(testCase.insertCarsErr)

			// when
			err := svc.PutCars(context.Background(), testCase.cars)

			// then
			assert.Equal(t, testCase.responseError, err)
		})
	}
}
