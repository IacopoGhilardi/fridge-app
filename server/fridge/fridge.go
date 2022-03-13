package fridge

import (
	"fridge-app/food"
	"time"
)

type Fridge struct {
	id        int
	food      []food.Food
	createdAt time.Time
}
