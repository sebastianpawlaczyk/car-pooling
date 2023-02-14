package healthCheck_api

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (h *HealthCheckHandler) GetStatus(c echo.Context) error {
	ctx := c.Request().Context()
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err := h.mongoClient.Ping(ctxTimeout, nil)
	if err != nil {
		msg := "mongo client is not alive!"
		log.Err(err).Msg(msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, "Service ready!")
}
