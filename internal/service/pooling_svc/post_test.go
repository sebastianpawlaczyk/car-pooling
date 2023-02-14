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

func TestPostJourney(t *testing.T) {
	testCases := []struct {
		name                      string
		journey                   *model.Journey
		getCarResponse            *model.Car
		getCarByAvailableSeatsErr error
		updateCarErr              error
		insertJourneyErr          error
		responseError             error
	}{
		{
			name: "ok",
			journey: &model.Journey{
				ID:     1,
				People: 4,
			},
			getCarResponse: &model.Car{
				ID:    2,
				Seats: 3,
			},
			getCarByAvailableSeatsErr: nil,
			updateCarErr:              nil,
			insertJourneyErr:          nil,
			responseError:             nil,
		},
		{
			name: "getCarByAvailableSeatsErr",
			journey: &model.Journey{
				ID:     1,
				People: 4,
			},
			getCarResponse:            nil,
			getCarByAvailableSeatsErr: fmt.Errorf("getCarByAvailableSeatsErr"),
			updateCarErr:              nil,
			insertJourneyErr:          nil,
			responseError:             fmt.Errorf("getCarByAvailableSeatsErr"),
		},
		{
			name: "updateCarErr",
			journey: &model.Journey{
				ID:     1,
				People: 4,
			},
			getCarResponse: &model.Car{
				ID:    2,
				Seats: 3,
			},
			getCarByAvailableSeatsErr: nil,
			updateCarErr:              fmt.Errorf("updateCarErr"),
			insertJourneyErr:          nil,
			responseError:             fmt.Errorf("updateCarErr"),
		},
		{
			name: "insertJourneyErr",
			journey: &model.Journey{
				ID:     1,
				People: 4,
			},
			getCarResponse: &model.Car{
				ID:    2,
				Seats: 3,
			},
			getCarByAvailableSeatsErr: nil,
			updateCarErr:              nil,
			insertJourneyErr:          fmt.Errorf("insertJourneyErr"),
			responseError:             fmt.Errorf("insertJourneyErr"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mockRepo := &mocks.PoolingRepository{}
			svc := NewPoolingService(mockRepo)

			mockRepo.On("GetCarByAvailableSeats", mock.Anything, mock.Anything).Return(testCase.getCarResponse, testCase.getCarByAvailableSeatsErr)
			mockRepo.On("UpdateCar", mock.Anything, mock.Anything).Return(testCase.updateCarErr)
			mockRepo.On("InsertJourney", mock.Anything, mock.Anything).Return(testCase.insertJourneyErr)

			// when
			err := svc.PostJourney(context.Background(), testCase.journey)

			// then
			assert.Equal(t, testCase.responseError, err)
		})
	}
}

func TestPostDropoff(t *testing.T) {
	testCases := []struct {
		name                 string
		journeyID            int
		getJourneyResponse   *model.Journey
		getJourneyByIDErr    error
		deleteJourneyByIDErr error
		getCarResponse       *model.Car
		getCarByIDErr        error
		updateNextJourneyErr error
		updateCarErr         error
		responseError        error
	}{
		{
			name:      "ok",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "3",
			},
			getJourneyByIDErr:    nil,
			deleteJourneyByIDErr: nil,
			getCarResponse: &model.Car{
				ID:    3,
				Seats: 4,
			},
			getCarByIDErr:        nil,
			updateNextJourneyErr: nil,
			updateCarErr:         nil,
			responseError:        nil,
		},
		{
			name:                 "getJourneyByIDErr",
			journeyID:            2,
			getJourneyResponse:   nil,
			getJourneyByIDErr:    fmt.Errorf("getJourneyByIDErr"),
			deleteJourneyByIDErr: nil,
			getCarResponse:       nil,
			getCarByIDErr:        nil,
			updateNextJourneyErr: nil,
			updateCarErr:         nil,
			responseError:        fmt.Errorf("getJourneyByIDErr"),
		},
		{
			name:      "deleteJourneyByIDErr",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "3",
			},
			getJourneyByIDErr:    nil,
			deleteJourneyByIDErr: fmt.Errorf("deleteJourneyByIDErr"),
			getCarResponse:       nil,
			getCarByIDErr:        nil,
			updateNextJourneyErr: nil,
			updateCarErr:         nil,
			responseError:        fmt.Errorf("deleteJourneyByIDErr"),
		},
		{
			name:      "getCarByIDErr",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "3",
			},
			getJourneyByIDErr:    nil,
			deleteJourneyByIDErr: nil,
			getCarResponse:       nil,
			getCarByIDErr:        fmt.Errorf("getCarByIDErr"),
			updateNextJourneyErr: nil,
			updateCarErr:         nil,
			responseError:        fmt.Errorf("getCarByIDErr"),
		},
		{
			name:      "car not found",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "3",
			},
			getJourneyByIDErr:    nil,
			deleteJourneyByIDErr: nil,
			getCarResponse:       nil,
			getCarByIDErr:        nil,
			updateNextJourneyErr: nil,
			updateCarErr:         nil,
			responseError:        fmt.Errorf("can not find assigned car"),
		},
		{
			name:      "updateNextJourneyErr",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "3",
			},
			getJourneyByIDErr:    nil,
			deleteJourneyByIDErr: nil,
			getCarResponse: &model.Car{
				ID:    3,
				Seats: 4,
			},
			getCarByIDErr:        nil,
			updateNextJourneyErr: fmt.Errorf("updateNextJourneyErr"),
			updateCarErr:         nil,
			responseError:        fmt.Errorf("updateNextJourneyErr"),
		},
		{
			name:      "updateCarErr",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "3",
			},
			getJourneyByIDErr:    nil,
			deleteJourneyByIDErr: nil,
			getCarResponse: &model.Car{
				ID:    3,
				Seats: 4,
			},
			getCarByIDErr:        nil,
			updateNextJourneyErr: nil,
			updateCarErr:         fmt.Errorf("updateCarErr"),
			responseError:        fmt.Errorf("updateCarErr"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mockRepo := &mocks.PoolingRepository{}
			svc := NewPoolingService(mockRepo)

			mockRepo.On("GetJourneyByID", mock.Anything, mock.Anything).Return(testCase.getJourneyResponse, testCase.getJourneyByIDErr)
			mockRepo.On("DeleteJourneyByID", mock.Anything, mock.Anything).Return(testCase.deleteJourneyByIDErr)
			mockRepo.On("GetCarByID", mock.Anything, mock.Anything).Return(testCase.getCarResponse, testCase.getCarByIDErr)
			mockRepo.On("UpdateNextPossibleJourney", mock.Anything, mock.Anything).Return(nil, testCase.updateNextJourneyErr)
			mockRepo.On("UpdateCar", mock.Anything, mock.Anything).Return(testCase.updateCarErr)

			// when
			journey, err := svc.PostDropoff(context.Background(), testCase.journeyID)

			// then
			assert.Equal(t, testCase.responseError, err)
			if err == nil {
				assert.Equal(t, testCase.getJourneyResponse, journey)
			}
		})
	}
}

func TestPostLocate(t *testing.T) {
	testCases := []struct {
		name                  string
		journeyID             int
		getJourneyResponse    *model.Journey
		getJourneyByIDErr     error
		getCarResponse        *model.Car
		getCarByIDErr         error
		getJourneysResponse   []*model.Journey
		getJourneysByCarIDErr error
		responseCar           *model.Car
		responseError         error
	}{
		{
			name:      "ok",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "1",
			},
			getJourneyByIDErr: nil,
			getCarResponse: &model.Car{
				ID:    1,
				Seats: 0,
			},
			getCarByIDErr: nil,
			getJourneysResponse: []*model.Journey{
				{
					ID:     2,
					People: 4,
					CarID:  "1",
				},
			},
			getJourneysByCarIDErr: nil,
			responseCar: &model.Car{
				ID:    1,
				Seats: 4,
			},
			responseError: nil,
		},
		{
			name:                  "ok",
			journeyID:             2,
			getJourneyResponse:    nil,
			getJourneyByIDErr:     fmt.Errorf("getJourneyByIDErr"),
			getCarResponse:        nil,
			getCarByIDErr:         nil,
			getJourneysResponse:   nil,
			getJourneysByCarIDErr: nil,
			responseCar:           nil,
			responseError:         fmt.Errorf("getJourneyByIDErr"),
		},
		{
			name:                  "journey not found",
			journeyID:             2,
			getJourneyResponse:    nil,
			getJourneyByIDErr:     nil,
			getCarResponse:        nil,
			getCarByIDErr:         nil,
			getJourneysResponse:   nil,
			getJourneysByCarIDErr: nil,
			responseCar:           nil,
			responseError:         CanNotLocateJourneyError,
		},
		{
			name:      "getCarByIDErr",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "1",
			},
			getJourneyByIDErr:     nil,
			getCarResponse:        nil,
			getCarByIDErr:         fmt.Errorf("getCarByIDErr"),
			getJourneysResponse:   nil,
			getJourneysByCarIDErr: nil,
			responseCar:           nil,
			responseError:         fmt.Errorf("getCarByIDErr"),
		},
		{
			name:      "car not found",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "1",
			},
			getJourneyByIDErr:     nil,
			getCarResponse:        nil,
			getCarByIDErr:         nil,
			getJourneysResponse:   nil,
			getJourneysByCarIDErr: nil,
			responseCar:           nil,
			responseError:         fmt.Errorf("can not find assigned car"),
		},
		{
			name:      "getJourneysByCarIDErr",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "1",
			},
			getJourneyByIDErr: nil,
			getCarResponse: &model.Car{
				ID:    1,
				Seats: 0,
			},
			getCarByIDErr: nil,
			getJourneysResponse: []*model.Journey{
				{
					ID:     2,
					People: 4,
					CarID:  "1",
				},
			},
			getJourneysByCarIDErr: fmt.Errorf("getJourneysByCarIDErr"),
			responseCar:           nil,
			responseError:         fmt.Errorf("error during finding journeys"),
		},
		{
			name:      "car not assigned",
			journeyID: 2,
			getJourneyResponse: &model.Journey{
				ID:     2,
				People: 4,
				CarID:  "",
			},
			getJourneyByIDErr:     nil,
			getCarResponse:        nil,
			getCarByIDErr:         nil,
			getJourneysResponse:   nil,
			getJourneysByCarIDErr: nil,
			responseCar:           nil,
			responseError:         nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mockRepo := &mocks.PoolingRepository{}
			svc := NewPoolingService(mockRepo)

			mockRepo.On("GetJourneyByID", mock.Anything, mock.Anything).Return(testCase.getJourneyResponse, testCase.getJourneyByIDErr)
			mockRepo.On("GetCarByID", mock.Anything, mock.Anything).Return(testCase.getCarResponse, testCase.getCarByIDErr)
			mockRepo.On("GetJourneysByCarID", mock.Anything, mock.Anything).Return(testCase.getJourneysResponse, testCase.getJourneysByCarIDErr)

			// when
			car, err := svc.PostLocate(context.Background(), testCase.journeyID)

			// then
			assert.Equal(t, testCase.responseError, err)
			assert.Equal(t, testCase.responseCar, car)
		})
	}
}
