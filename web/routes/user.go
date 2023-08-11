package routes

import (
	bookController "eventsourcing/web/controllers"

	"github.com/gin-gonic/gin"
)

var prefix string = "books"

func UserRoutes(router *gin.RouterGroup) {
	router.GET(prefix, bookController.CreateUser)
}
