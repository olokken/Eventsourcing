package user

import (
	"eventsourcing/application"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

//Applies the event to the user object

type IUserEvent interface {
	Apply(u *User)
	IsEvent()
}

type UserCreated struct {
	application.Event
	ObjectId    string
	Name        string
	Email       string
	PhoneNumber string
}

func (e UserCreated) IsEvent() {}

func (e UserCreated) Apply(u *User) {
	u.ObjectId = e.ObjectId
	u.Name = e.Name
	u.Email = e.Email
	u.PhoneNumber = e.PhoneNumber
}

type NameChanged struct {
	application.Event
	Name     string
	ObjectId string
}

func (e NameChanged) IsEvent() {}

func (e NameChanged) Apply(u *User) {
	u.Name = e.Name
}

func CreateEventFromBson(data bson.M) IUserEvent {
	event := data["event"].(bson.M)
	eventType := event["eventtype"].(int32)
	fmt.Print("Vamos")

	bsonBytes, err := bson.Marshal(data)

	if err != nil {
		fmt.Errorf("Failed in CreateEventFromBson")
	}

	switch eventType {
	case 0:
		var event UserCreated
		bson.Unmarshal(bsonBytes, &event)
		return event
	case 1:
		var event NameChanged
		bson.Unmarshal(bsonBytes, &event)
		return event
	default:
		return nil
	}
}
