package user

import "eventsourcing/application"

type User struct {
	application.Aggregate
	UserId      string
	Name        string
	Email       string
	PhoneNumber string
	Dockets     []int32
}

func (u *User) ApplyEvent(e IUserEvent) {
	u.AddEvent(e)
	e.Apply(u)
}
