package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/car-pooling/internal/cfg"
	"github.com/car-pooling/internal/endpoint/healthCheck_api"
	"github.com/car-pooling/internal/endpoint/pooling_api"
	"github.com/car-pooling/internal/repository/pooling_repo"
	"github.com/car-pooling/internal/server"
	"github.com/car-pooling/internal/service/pooling_svc"
)

func main() {
	config, err := cfg.LoadConfig()
	if err != nil {
		panic(err)
	}

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	carRepository := pooling_repo.NewPoolingRepository(client, config)

	poolingService := pooling_svc.NewPoolingService(carRepository)

	healthCheckHandler := healthCheck_api.NewHealthCheckHandler(client)
	poolingHandler := pooling_api.NewPoolingHandler(poolingService)

	httpServer := server.NewHttpServer(config, healthCheckHandler, poolingHandler)

	err = httpServer.Start()
	if err != nil {
		panic(err)
	}
}
