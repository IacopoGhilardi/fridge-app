package models

import (
	"fridge-app/food"
	"time"
)

type Fridge struct {
	Id         int         `bson:"_id,omitempty"`
	Food       []food.Food `bson:food`
	Created_at time.Time   `bson:created_at,omitempty`
}
