package pooling_repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/car-pooling/internal/cfg"
	"github.com/car-pooling/internal/model"
)

type PoolingRepository interface {
	InsertCars(ctx context.Context, cars []*model.Car) error
	DeleteCars(ctx context.Context) error
	UpdateCar(ctx context.Context, updatedCar *model.Car) error
	GetCarByAvailableSeats(ctx context.Context, seats int) (*model.Car, error)
	GetCarByID(ctx context.Context, carID int) (*model.Car, error)

	InsertJourney(ctx context.Context, journey *model.Journey) error
	DeleteJourneys(ctx context.Context) error
	DeleteJourneyByID(ctx context.Context, journeyID int) error
	GetJourneyByID(ctx context.Context, journeyID int) (*model.Journey, error)
	GetJourneysByCarID(ctx context.Context, carID string) ([]*model.Journey, error)
	UpdateNextPossibleJourney(ctx context.Context, car *model.Car) (*model.Journey, error)
}

type poolingRepository struct {
	client *mongo.Client
	config *cfg.Config
}

func NewPoolingRepository(client *mongo.Client, config *cfg.Config) PoolingRepository {
	return &poolingRepository{
		client: client,
		config: config,
	}
}

func (repo *poolingRepository) getDatabase() *mongo.Database {
	return repo.client.Database(repo.config.DatabaseName)
}
