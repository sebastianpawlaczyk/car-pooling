package pooling_api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/car-pooling/internal/model"
	"github.com/car-pooling/internal/service/pooling_svc"
)

func (h *PoolingHandler) PutCars(c echo.Context) error {
	var body []*model.Car
	if err := c.Bind(&body); err != nil {
		msg := "could not process provided JSON; verify its structure"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := validCars(body); err != nil {
		msg := "car list is not valid"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	err := h.poolingService.PutCars(ctx, body)
	if err != nil {
		msg := "error during loading new car list"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (h *PoolingHandler) PostJourney(c echo.Context) error {
	var body *model.Journey
	if err := c.Bind(&body); err != nil {
		msg := "could not process provided JSON; verify its structure"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := validJourney(body); err != nil {
		msg := "journey is not valid"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	err := h.poolingService.PostJourney(ctx, body)
	if err != nil {
		msg := "error during creating new journey"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (h *PoolingHandler) PostDropoff(c echo.Context) error {
	journeyID := c.FormValue("ID")

	id, err := strconv.Atoi(journeyID)
	if err != nil {
		msg := "error during parse int"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	journey, err := h.poolingService.PostDropoff(ctx, id)
	if err != nil {
		msg := "error during dropoff journey"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	if journey == nil {
		msg := "group not found"
		log.Info().Msg(msg)
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}

func (h *PoolingHandler) PostLocate(c echo.Context) error {
	journeyID := c.FormValue("ID")

	id, err := strconv.Atoi(journeyID)
	if err != nil {
		msg := "error during parse int"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	car, err := h.poolingService.PostLocate(ctx, id)
	if err != nil {
		if err == pooling_svc.CanNotLocateJourneyError {
			msg := "group not found"
			log.Info().Msg(msg)
			return c.NoContent(http.StatusNotFound)
		}
		msg := "error during locate journey"
		log.Err(err).Msg(msg)
		return c.NoContent(http.StatusBadRequest)
	}

	if car == nil {
		msg := "group is waiting to be assigned to a car"
		log.Info().Msg(msg)
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, car)
}
