package storages

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongo(uri, name string) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return &Mongo{
		Client:   client,
		Database: client.Database(name),
	}, nil
}

func (m *Mongo) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.Client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Failed to disconnect MongoDB client: %v", err)
	}

	log.Println("Disconnected MongoDB client")
}

//func (m *Mongo) Collection(name string) *mongo.Collection {
//	return m.Database.Collection(name)
//}
