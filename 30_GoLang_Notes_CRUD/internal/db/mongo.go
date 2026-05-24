package db

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {
	// prevents app from freezing during startup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create client options
	clientOpts := options.Client().ApplyURI(cfg.MongoURI)

	// connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect MongoDB: %v", err)
	}

	// ping MongoDB to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// select database
	database := client.Database(cfg.MongoDB)

	return client, database, nil
}

func Disconnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Disconnect(ctx)
}