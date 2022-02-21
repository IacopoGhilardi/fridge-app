package main

import (
	"fmt"
	"io/ioutil"
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
		bodyReq, _ := ioutil.ReadAll(c.Request.Body)
		println(string(body))
		foodParam := c.Param("food")
		println(string(foodParam))
		var newFood food.Food = food.NewFood("banana", 33)

		c.String(http.StatusOK, bodyReq, newFood)
	})

	banana := food.NewFood("banana", 33)
	fmt.Println("banaan", banana)

	router.Run("localhost:3000")

	// fmt.Println("Connected to MongoDB!")
}
