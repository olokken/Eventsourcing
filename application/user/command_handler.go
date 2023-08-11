package user

type CommandHandler struct {
	Repository string
}

func (ch *CommandHandler) HandleCommand(c IUserCommand) {
	c.Handle(ch)
}
