package routes

import (
	"fridge-app/controllers"

	"github.com/gin-gonic/gin"
)

func FridgeRoute(router *gin.Engine) {

	// router.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/")
	// })

	router.POST("/fridge/new", controllers.CreateFridge())

	// router.GET("/fridge", func(c *gin.Context) {
	// 	var fridges []fridge.Fridge
	// 	cursor, err := fridgesCollection.Find(ctx, bson.M{})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if err = cursor.All(ctx, &fridges); err != nil {
	// 		panic(err)
	// 	}
	// 	c.JSON(200, gin.H{"foodList": fridges[0].Food})
	// })

	// router.POST("/fridge/upsert/:id", func(context *gin.Context) {
	// 	// body, err := ioutil.ReadAll(context.Request.Body)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }

	// 	// fridgeId := context.Request.Header["id"]
	// 	fridgeId := 2

	// 	var foodArray []food.Food = []food.Food{
	// 		{Name: "Banana", Image: "banana.jpg", Quantity: 2},
	// 		{Name: "Apple", Image: "apple.jpg", Quantity: 2},
	// 	}

	// 	newFridge := fridge.Fridge{
	// 		Id:         fridgeId,
	// 		Food:       foodArray,
	// 		Created_at: time.Now(),
	// 	}

	// 	insertResult, err := fridgesCollection.InsertOne(ctx, newFridge)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	context.JSON(200, gin.H{"newFridge": newFridge, "result": insertResult})
	// })
}
