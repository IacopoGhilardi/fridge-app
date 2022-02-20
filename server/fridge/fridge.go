package fridge

import (
	"fmt"
	"fridge-app/food"
)

type Fridge struct {
	id   int
	food []food.Food
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
