package user

type IUserCommand interface {
	Handle(c *CommandHandler)
}

//Handles the businesslogic, creates an event if it should be created and applies it

type CreateUser struct {
	UserId      string
	Name        string
	PhoneNumber string
	Email       string
}

func (c CreateUser) Handle(handler *CommandHandler) {
	var u User = User{}

	//Handle business logic
	var event UserCreated = UserCreated{Name: c.Name, Email: c.Email, PhoneNumber: c.PhoneNumber, UserId: c.UserId}
	u.ApplyEvent(event)
}
