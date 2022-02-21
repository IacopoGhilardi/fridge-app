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

func PrintSomething() {
	fmt.Println("speriamo funzioni cazzo")
}

func NewFridge(id int, food []food.Food) *Fridge {
	newFridge := new(Fridge)
	newFridge.id = id
	newFridge.food = food
	return newFridge
}
