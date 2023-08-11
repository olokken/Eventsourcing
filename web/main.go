package main

import (
	"eventsourcing/web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		routes.UserRoutes(v1)
	}

	router.Run(":8080")
}
