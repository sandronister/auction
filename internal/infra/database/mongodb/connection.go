package mongodb

import (
	"context"

	"github.com/sandronister/auction-go/internal/infra/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewConnection(uri string,database string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error("Error connecting to MongoDB", err)
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		logger.Error("Error pinging MongoDB", err)
		return nil, err
	}

	return client.Database(database), nil
}