package pooling_api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/car-pooling/internal/model"
	"github.com/car-pooling/internal/service/pooling_svc"
	mocks "github.com/car-pooling/mocks/mocks_svc"
)

func TestPutCars(t *testing.T) {
	testCases := []struct {
		name         string
		svcError     error
		wrongRequest bool
		responseCode int
		finalError   error
	}{
		{
			name:         "ok",
			svcError:     nil,
			responseCode: http.StatusOK,
			finalError:   nil,
		},
		{
			name:         "wrong request",
			svcError:     nil,
			wrongRequest: true,
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
		{
			name:         "svc error",
			svcError:     fmt.Errorf("some svc error"),
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mock_svc := &mocks.PoolingInterface{}
			handler := NewPoolingHandler(mock_svc)
			rr := httptest.NewRecorder()

			jsonData := []byte(`[
				{
					"id": 1,
					"seats": 4
				},
				{
					"id": 2,
					"seats": 6
				}
			]`)

			if testCase.wrongRequest {
				jsonData = []byte(`[`)
			}

			req := httptest.NewRequest(http.MethodPut, CarsPath, bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			mock_svc.On("PutCars", mock.Anything, mock.Anything).Return(testCase.svcError)

			// when
			err := handler.PutCars(echo.New().NewContext(req, rr))

			// then
			assert.Equal(t, testCase.responseCode, rr.Code)
			assert.Equal(t, testCase.finalError, err)
		})
	}
}

func TestPostJourney(t *testing.T) {
	testCases := []struct {
		name         string
		svcError     error
		wrongRequest bool
		responseCode int
		finalError   error
	}{
		{
			name:         "ok",
			svcError:     nil,
			responseCode: http.StatusOK,
			finalError:   nil,
		},
		{
			name:         "wrong request",
			svcError:     nil,
			wrongRequest: true,
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
		{
			name:         "svc error",
			svcError:     fmt.Errorf("some svc error"),
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mock_svc := &mocks.PoolingInterface{}
			handler := NewPoolingHandler(mock_svc)
			rr := httptest.NewRecorder()

			jsonData := []byte(`{
				"id": 1,
				"people": 5
			}`)

			if testCase.wrongRequest {
				jsonData = []byte(`[`)
			}

			req := httptest.NewRequest(http.MethodPut, CarsPath, bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			mock_svc.On("PostJourney", mock.Anything, mock.Anything).Return(testCase.svcError)

			// when
			err := handler.PostJourney(echo.New().NewContext(req, rr))

			// then
			assert.Equal(t, testCase.responseCode, rr.Code)
			assert.Equal(t, testCase.finalError, err)
		})
	}
}

func TestPostDropoff(t *testing.T) {
	testCases := []struct {
		name         string
		svcResponse  *model.Journey
		svcError     error
		wrongRequest bool
		responseCode int
		finalError   error
	}{
		{
			name: "ok",
			svcResponse: &model.Journey{
				ID:     2,
				People: 5,
				CarID:  "1",
			},
			svcError:     nil,
			responseCode: http.StatusOK,
			finalError:   nil,
		},
		{
			name:         "bad request",
			svcResponse:  nil,
			svcError:     nil,
			wrongRequest: true,
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
		{
			name: "svc error",
			svcResponse: &model.Journey{
				ID:     2,
				People: 5,
				CarID:  "1",
			},
			svcError:     fmt.Errorf("some svc error"),
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
		{
			name:         "journey not found",
			svcResponse:  nil,
			svcError:     nil,
			responseCode: http.StatusNotFound,
			finalError:   nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mock_svc := &mocks.PoolingInterface{}
			handler := NewPoolingHandler(mock_svc)
			rr := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodPut, CarsPath, nil)
			req.Form = url.Values{}
			if testCase.wrongRequest {
				req.Form.Add("ID", "A")
			} else {
				req.Form.Add("ID", "2")
			}

			mock_svc.On("PostDropoff", mock.Anything, mock.Anything).Return(testCase.svcResponse, testCase.svcError)

			// when
			err := handler.PostDropoff(echo.New().NewContext(req, rr))

			// then
			assert.Equal(t, testCase.responseCode, rr.Code)
			assert.Equal(t, testCase.finalError, err)
		})
	}
}

func TestPostLocate(t *testing.T) {
	testCases := []struct {
		name         string
		svcResponse  *model.Car
		svcError     error
		wrongRequest bool
		responseCode int
		finalError   error
	}{
		{
			name: "ok",
			svcResponse: &model.Car{
				ID:    1,
				Seats: 5,
			},
			svcError:     nil,
			responseCode: http.StatusOK,
			finalError:   nil,
		},
		{
			name:         "bad request",
			svcResponse:  nil,
			svcError:     nil,
			wrongRequest: true,
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
		{
			name:         "group not found",
			svcResponse:  nil,
			svcError:     pooling_svc.CanNotLocateJourneyError,
			responseCode: http.StatusNotFound,
			finalError:   nil,
		},
		{
			name:         "svc error",
			svcResponse:  nil,
			svcError:     fmt.Errorf("some svc error"),
			responseCode: http.StatusBadRequest,
			finalError:   nil,
		},
		{
			name:         "no car assigned",
			svcResponse:  nil,
			svcError:     nil,
			responseCode: http.StatusNoContent,
			finalError:   nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// given
			mock_svc := &mocks.PoolingInterface{}
			handler := NewPoolingHandler(mock_svc)
			rr := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodPut, CarsPath, nil)
			req.Form = url.Values{}
			if testCase.wrongRequest {
				req.Form.Add("ID", "A")
			} else {
				req.Form.Add("ID", "2")
			}

			mock_svc.On("PostLocate", mock.Anything, mock.Anything).Return(testCase.svcResponse, testCase.svcError)

			// when
			err := handler.PostLocate(echo.New().NewContext(req, rr))

			// then
			assert.Equal(t, testCase.responseCode, rr.Code)
			assert.Equal(t, testCase.finalError, err)
		})
	}
}
