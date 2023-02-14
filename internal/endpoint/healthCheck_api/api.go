package healthCheck_api

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
	mongoClient *mongo.Client
}

func NewHealthCheckHandler(mongoClient *mongo.Client) *HealthCheckHandler {
	return &HealthCheckHandler{
		mongoClient: mongoClient,
	}
}

func (h *HealthCheckHandler) Register(router *echo.Echo) {
	group := router.Group(EmptyPrefix)

	group.GET(StatusPath, h.GetStatus)
}
