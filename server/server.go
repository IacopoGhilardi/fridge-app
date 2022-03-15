package main

import (
	"log"
	"net/http"
	"time"

	"fridge-app/food"
	"fridge-app/fridge"
	"fridge-app/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	name    string
	address string
}

func main() {

	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := setup.Connect("mongodb://localhost:27017/the-fridge")
	if err != nil {
		panic(err)
	}

	database := client.Database("the-fridge")
	fridgesCollection := database.Collection("fridges")

	// Free the resource when main function is  returned
	defer setup.Close(client, ctx, cancel)
	setup.Ping(client, ctx)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/")
	})

	router.GET("/fridge", func(c *gin.Context) {
		var fridges []fridge.Fridge
		cursor, err := fridgesCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		if err = cursor.All(ctx, &fridges); err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{"foodList": fridges[0].Food})
	})

	router.POST("/fridge/upsert/:id", func(context *gin.Context) {
		// body, err := ioutil.ReadAll(context.Request.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fridgeId := context.Request.Header["id"]
		fridgeId := 2

		var foodArray []food.Food = []food.Food{
			{Name: "Banana", Image: "banana.jpg", Quantity: 2},
			{Name: "Apple", Image: "apple.jpg", Quantity: 2},
		}

		newFridge := fridge.Fridge{
			Id:         fridgeId,
			Food:       foodArray,
			Created_at: time.Now(),
		}

		insertResult, err := fridgesCollection.InsertOne(ctx, newFridge)

		if err != nil {
			log.Fatal(err)
		}

		context.JSON(200, gin.H{"newFridge": newFridge, "result": insertResult})
	})

	router.Run("localhost:3000")
}
