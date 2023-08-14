package user

import (
	"eventsourcing/repository"
	"fmt"
)

type IQueryHandler interface {
	GetUser(id string) User
}

type QueryHandler struct {
	Repository repository.IEventRepository //Chaning this to a Read Database later
}

func (q QueryHandler) GetUser(id string) User {
	var user User = User{}
	events := q.Repository.GetByObjectId(id)

	for _, event := range events {
		var userEvent IUserEvent = CreateEventFromBson(event)
		user.ApplyEvent(userEvent)
	}

	fmt.Print(events)
	return user
}
