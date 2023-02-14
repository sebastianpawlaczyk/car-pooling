package pooling_api

import (
	"github.com/labstack/echo/v4"

	"github.com/car-pooling/internal/service/pooling_svc"
)

type PoolingHandler struct {
	poolingService pooling_svc.PoolingInterface
}

func NewPoolingHandler(poolingService pooling_svc.PoolingInterface) *PoolingHandler {
	return &PoolingHandler{
		poolingService: poolingService,
	}
}

func (h *PoolingHandler) Register(router *echo.Echo) {
	group := router.Group(EmptyPrefix)

	group.PUT(CarsPath, h.PutCars)
	group.POST(JourneyPath, h.PostJourney)
	group.POST(DropoffPath, h.PostDropoff)
	group.POST(LocatePath, h.PostLocate)
}
