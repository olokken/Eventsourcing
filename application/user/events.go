package user

import "eventsourcing/application"

//Applies the event to the user object

type IUserEvent interface {
	Apply(u *User)
	IsEvent()
}

type UserCreated struct {
	application.Event
	UserId      string
	Name        string
	Email       string
	PhoneNumber string
}

func (e UserCreated) IsEvent() {}

func (e UserCreated) Apply(u *User) {
	u.UserId = e.UserId
	u.Name = e.Name
	u.Email = e.Email
	u.PhoneNumber = e.PhoneNumber
}

type NameChanged struct {
	application.Event
	name string
}

func (e NameChanged) isEvent() {}

func (e NameChanged) Apply(u *User) {
	u.Name = e.name
}
