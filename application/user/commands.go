package user

import "eventsourcing/application"

type IUserCommand interface {
	Handle(c *CommandHandler)
}

//Handles the businesslogic, creates an event if it should be created and applies it

type CreateUser struct {
	ObjectId    string
	Name        string
	PhoneNumber string
	Email       string
}

func (c CreateUser) Handle(handler *CommandHandler) {
	var u User = User{}
	event := application.CreateEvent(application.User, application.UserCreated, 0)
	var createUserEvent UserCreated = UserCreated{Name: c.Name, Email: c.Email, PhoneNumber: c.PhoneNumber, ObjectId: c.ObjectId, Event: event}
	u.ApplyEvent(createUserEvent)
	handler.Repository.Save(u.Events[0])
}

type ChangeName struct {
	ObjectId string
	Name     string
}

func (c ChangeName) Handle(handler *CommandHandler) {
	var u User = User{Name: c.Name, ObjectId: c.ObjectId} //Need to get User to find version and do businesslogic
	event := application.CreateEvent(application.User, application.NameChanged, 1)
	changedNameEvent := NameChanged{Name: c.Name, ObjectId: c.ObjectId, Event: event}
	u.ApplyEvent(changedNameEvent)
	handler.Repository.Save(u.Events[0])
}
