package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"fridge-app/food"
	"fridge-app/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := setup.Connect("mongodb://localhost:27017/the-fridge")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function is  returned
	defer setup.Close(client, ctx, cancel)
	setup.Ping(client, ctx)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/")
	})

	router.GET("/fridge", func(c *gin.Context) {
		myMap := map[string]int{"prova": 12, "altra_prova": 13}
		c.JSON(200, myMap)
	})

	router.POST("/fridge/upsert/:id", func(context *gin.Context) {
		body, err := ioutil.ReadAll(context.Request.Body)
		bodyRequest := context.Request.Body

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(bodyRequest)
		fmt.Print(body)

		foodParam := context.Param("food")
		println(string(foodParam))
		var newFood = new(food.Food)
		newFood.Init("banana", 33)

		context.JSON(http.StatusOK, gin.H{"newFood": newFood})
		// c.String(http.StatusOK, bodyRequest, newFood)
	})

	router.Run("localhost:3000")
	// fmt.Println("Connected to MongoDB!")
}
