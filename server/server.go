package main

import (
	"fmt"
	"net/http"

	"fridge-app/food"
	"fridge-app/fridge"

	"github.com/gin-gonic/gin"
)

func main() {

	var router = gin.Default()

	fridge.PrintSomething()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/")
	})

	router.POST("/fridge/upsert/:id", func(c *gin.Context) {
		food := c.Param("food")
		newFood := food.NewFood(food.name, food.qty)

		c.String(http.StatusOK, newFood)
	})

	banana := food.NewFood("banana", 33)
	fmt.Println("banaan", banana)

	router.Run("localhost:3000")

	// fmt.Println("Connected to MongoDB!")
}
