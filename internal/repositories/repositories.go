package repositories

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/models"
	"awesomeProject/internal/storages"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Config     *config.Config
	collection *mongo.Collection
}

func NewRepository(db *storages.Mongo, cfg *config.Config) *Repository {
	return &Repository{
		Config:     cfg,
		collection: db.Database.Collection("shorten"),
	}
}

func (r *Repository) CreateByLong(ctx context.Context, shorten models.Shorten) (*models.Shorten, error) {
	result, err := r.collection.InsertOne(ctx, shorten)
	if err != nil {
		return nil, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to cast InsertedID")
	}

	shorten.ID = id
	return &shorten, nil
}

func (r *Repository) FindLongByShorten(ctx context.Context, shortenUrl string) (*models.Shorten, error) {
	var result models.Shorten

	err := r.collection.FindOne(ctx, bson.M{"short": shortenUrl}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
