package controllers

import (
	"eventsourcing/application/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	CreateUser(c *gin.Context)
	ChangeName(c *gin.Context)
	GetUser(c *gin.Context)
}

type UserController struct {
	CommandHandler user.CommandHandler
	QueryHander    user.IQueryHandler
}

func (controller UserController) CreateUser(c *gin.Context) {
	var userCommand user.CreateUser

	if err := c.ShouldBindJSON(&userCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.CommandHandler.HandleCommand(userCommand)
	c.JSON(200, userCommand)
}

func (controller UserController) ChangeName(c *gin.Context) {
	var changeNameCommand user.ChangeName

	if err := c.ShouldBindJSON(&changeNameCommand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.CommandHandler.HandleCommand(changeNameCommand)
	c.JSON(200, changeNameCommand)
}

func (controller UserController) GetUser(c *gin.Context) {
	user := controller.QueryHander.GetUser("123")
	c.JSON(200, user)
}
