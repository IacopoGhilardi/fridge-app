package controllers

import (
	"context"
	"fridge-app/configs"
	"fridge-app/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var fridgeCollection *mongo.Collection = configs.GetCollection(configs.Database, "fridges")
var validate = validator.New()

func CreateFridge() gin.HandlerFunc {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return func(context *gin.Context) {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var fridge models.Fridge
		defer cancel()

		if err := context.BindJSON(&fridge); err != nil {
			context.JSON(500, gin.H{"error": err})
			log.Fatal(err)
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&fridge); validationErr != nil {
			context.JSON(500, gin.H{"validate_error": validationErr})
			return
		}

		newFridge := models.Fridge{
			Id:   fridge.Id,
			Food: fridge.Food,
		}

		result, err := fridgeCollection.InsertOne(ctx, newFridge)

		if err != nil {
			context.JSON(500, gin.H{"error": err})
		}

		context.JSON(200, gin.H{"data": result})

	}
}
