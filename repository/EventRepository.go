package repository

import (
	"context"
	"eventsourcing/application"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IEventRepository interface {
	Save(event application.IEvent)
	GetByObjectId(id string) []bson.M
}

type EventRepository struct {
	Collection *mongo.Collection
}

func (e EventRepository) Save(event application.IEvent) {
	result, err := e.Collection.InsertOne(context.TODO(), event)
	fmt.Print(result)
	fmt.Print(err)
}

func (e EventRepository) GetByObjectId(id string) []bson.M {
	var results []bson.M

	filter := bson.D{{Key: "objectid", Value: id}}
	cursor, err := e.Collection.Find(context.TODO(), filter)

	if err != nil {
		// Handle error
		return results
	}

	defer cursor.Close(context.TODO()) // Close the cursor when done

	for cursor.Next(context.TODO()) {
		//Here check which event it is, then map the cursor to the eventType
		var result bson.M

		err := cursor.Decode(&result)

		if err != nil {
			continue
		}

		results = append(results, result) // Append the event to the slice
	}

	return results
}
