package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"fridge-app/food"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func getAllFridges(client *mongo.Client, ctx context.Context) {

}

func main() {

	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017/the-fridge")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function is  returned
	defer close(client, ctx, cancel)
	ping(client, ctx)

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
