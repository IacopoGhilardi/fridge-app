package main

import (
	"fmt"
	"net/http"

	"fridge-app/food"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/")
	})

	router.GET("/fridge", func(c *gin.Context) {
		myMap := map[string]int{"prova": 12, "altra_prova": 13}
		c.JSON(200, myMap)
	})

	router.POST("/fridge/upsert/:id", func(c *gin.Context) {
		bodyRequest := c.Request.Body

		fmt.Print(bodyRequest)
		// bodReq, _ := ioutil.ReadAll(c.Request.Body)
		// println(string(bodyReq))
		foodParam := c.Param("food")
		println(string(foodParam))
		var newFood = new(food.Food)
		newFood.Init("banana", 33)

		c.JSON(200, bodyRequest)
		// c.String(http.StatusOK, bodyRequest, newFood)
	})

	router.Run("localhost:3000")
	// fmt.Println("Connected to MongoDB!")
}
