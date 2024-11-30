package mongodb;

import (
	"go.mongodb.org/mongo-driver/mongo"
	
)

func NewMongoDbConnection(Config *enviroment.Config) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}