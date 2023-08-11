package controllers

import (
	"eventsourcing/application/user"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	handler := user.CommandHandler{Repository: "Repository"}

	var command user.CreateUser = user.CreateUser{UserId: "123", Name: "Navn", Email: "ole@mail.com", PhoneNumber: "12345"}
	handler.HandleCommand(command)

	c.JSON(200, "vamos")
}
