package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	DB     *mongo.Database
	client *mongo.Client
}

func NewWithConfig(ctx context.Context, c *Config) (*DB, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(c.ConnectionString()).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect with mongodb: %w", err)
	}

	return &DB{
		DB:     client.Database("admin"),
		client: client,
	}, nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
