package fridge

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fridge struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Food       []string           `bson:food`
	Created_at time.Time          `bson:created_at,omitempty`
}
