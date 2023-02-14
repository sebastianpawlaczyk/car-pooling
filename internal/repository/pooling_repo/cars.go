package pooling_repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/car-pooling/internal/model"
	"github.com/car-pooling/internal/repository"
)

func (repo *poolingRepository) DeleteCars(ctx context.Context) error {
	err := repo.getDatabase().Collection(repository.CarsCollection).Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *poolingRepository) InsertCars(ctx context.Context, cars []*model.Car) error {
	carsCollection := repo.getDatabase().Collection(repository.CarsCollection)
	toInsert := []interface{}{}
	for _, car := range cars {
		toInsert = append(toInsert, car)
	}

	_, err := carsCollection.InsertMany(ctx, toInsert)
	if err != nil {
		return err
	}

	return nil
}

func (repo *poolingRepository) UpdateCar(ctx context.Context, updatedCar *model.Car) error {
	carsCollection := repo.getDatabase().Collection(repository.CarsCollection)

	filter := bson.D{{"_id", updatedCar.ID}}
	update := bson.D{{"$set", updatedCar}}

	_, err := carsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *poolingRepository) GetCarByAvailableSeats(ctx context.Context, seats int) (*model.Car, error) {
	carsCollection := repo.getDatabase().Collection(repository.CarsCollection)

	setSort := options.Find().SetSort(bson.D{{"seats", 1}}).SetLimit(1)

	result, err := carsCollection.Find(ctx, bson.M{"seats": bson.M{"$gte": seats}}, setSort)
	if err != nil {
		return nil, err
	}
	cars := []model.Car{}
	if err := result.All(ctx, &cars); err != nil {
		return nil, err
	}

	if len(cars) == 0 {
		return nil, nil
	}

	return &cars[0], nil
}

func (repo *poolingRepository) GetCarByID(ctx context.Context, carID int) (*model.Car, error) {
	carsCollection := repo.getDatabase().Collection(repository.CarsCollection)

	filter := bson.D{{"_id", carID}}

	result := carsCollection.FindOne(ctx, filter)

	car := model.Car{}
	if err := result.Decode(&car); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &car, nil
}
