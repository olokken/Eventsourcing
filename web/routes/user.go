package routes

import (
	"eventsourcing/web/controllers"

	"github.com/gin-gonic/gin"
)

var prefix string = "user"

func UserRoutes(router *gin.RouterGroup, controller controllers.IUserController) {
	router.POST(prefix, controller.CreateUser)
	router.PATCH(prefix, controller.ChangeName)
	router.GET(prefix, controller.GetUser)
}
