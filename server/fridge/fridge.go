package fridge

import (
	"fmt"
	"fridge-app/food"
	"time"
)

type Fridge struct {
	id        int
	food      []food.Food
	createdAt time.Time
}

func (newFridge *Fridge) Init((id int, food []) {
	newFridge := new(Fridge)
	newFridge.id = id
	newFridge.food = food
}
