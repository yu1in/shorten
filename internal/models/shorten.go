package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Shorten struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Long      string             `json:"long" bson:"long"`
	Short     string             `json:"short" bson:"short"`
	IsActive  bool               `json:"isActive" bson:"isActive"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
