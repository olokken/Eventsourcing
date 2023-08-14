package main

import (
	"context"
	"eventsourcing/application/user"
	"eventsourcing/repository"
	"eventsourcing/web/controllers"
	"eventsourcing/web/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := gin.Default()

	collection, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	repository := repository.EventRepository{Collection: collection}

	var userHandler user.CommandHandler = user.CommandHandler{Repository: repository}
	var userQueryHandler user.IQueryHandler = user.QueryHandler{Repository: repository}

	var userController controllers.IUserController = controllers.UserController{CommandHandler: userHandler, QueryHander: userQueryHandler}

	v1 := router.Group("api/v1")
	{
		routes.UserRoutes(v1, userController)
	}

	router.Run(":8080")
}

func connectToDatabase() (*mongo.Collection, error) {

	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	ctx := context.Background() // You can replace this with a more specific context

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	collection := client.Database("events").Collection("events")
	return collection, nil

}
