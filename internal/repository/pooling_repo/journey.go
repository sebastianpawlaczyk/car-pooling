package pooling_repo

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/car-pooling/internal/model"
	"github.com/car-pooling/internal/repository"
)

func (repo *poolingRepository) InsertJourney(ctx context.Context, journey *model.Journey) error {
	journeyCollection := repo.getDatabase().Collection(repository.JourneyCollection)

	_, err := journeyCollection.InsertOne(ctx, journey)
	if err != nil {
		return err
	}

	return nil
}

func (repo *poolingRepository) DeleteJourneys(ctx context.Context) error {
	err := repo.getDatabase().Collection(repository.JourneyCollection).Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *poolingRepository) DeleteJourneyByID(ctx context.Context, journeyID int) error {
	journeyCollection := repo.getDatabase().Collection(repository.JourneyCollection)

	filter := bson.D{{"_id", journeyID}}
	_, err := journeyCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (repo *poolingRepository) GetJourneyByID(ctx context.Context, journeyID int) (*model.Journey, error) {
	journeyCollection := repo.getDatabase().Collection(repository.JourneyCollection)

	filter := bson.D{{"_id", journeyID}}
	result := journeyCollection.FindOne(ctx, filter)

	journey := model.Journey{}
	if err := result.Decode(&journey); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &journey, nil
}

func (repo *poolingRepository) GetJourneysByCarID(ctx context.Context, carID string) ([]*model.Journey, error) {
	journeyCollection := repo.getDatabase().Collection(repository.JourneyCollection)

	filter := bson.D{{"carID", carID}}
	result, err := journeyCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	journeys := []*model.Journey{}
	if err := result.All(ctx, &journeys); err != nil {
		return nil, err
	}

	return journeys, nil
}

func (repo *poolingRepository) UpdateNextPossibleJourney(ctx context.Context, car *model.Car) (*model.Journey, error) {
	journeyCollection := repo.getDatabase().Collection(repository.JourneyCollection)

	filter := bson.M{
		"people": bson.M{"$lte": car.Seats},
		"carID":  bson.M{"$eq": ""},
	}
	setSort := options.FindOneAndUpdate().SetSort(bson.D{{"createdOn", 1}})
	update := bson.M{"$set": bson.M{"carID": strconv.Itoa(car.ID)}}

	result := journeyCollection.FindOneAndUpdate(ctx, filter, update, setSort)

	journey := model.Journey{}
	if err := result.Decode(&journey); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &journey, nil
}
