package app

import (
	"context"
	"fmt"
	"time"

	"go-auth/internal/config"
	"go-auth/internal/db"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config config.Config
	MongoClient *mongo.Client
	DB *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	//load env first, then connect to mongo and return the app struct
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	mongoClient, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return &App{
		Config: cfg,
		MongoClient: mongoClient.Client,
		DB: mongoClient.DB,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB: %w", err)
	}

	return nil
}
