package user

import "eventsourcing/repository"

type CommandHandler struct {
	Repository repository.IEventRepository
}

func (ch *CommandHandler) HandleCommand(c IUserCommand) {
	c.Handle(ch)
}
