package main

import (
	"github.com/mutemaniac/steel/functions"
	"github.com/mutemaniac/steel/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.POST("/post", createRoute)
	router.Run(":8080")
}

func createRoute(c *gin.Context) {
	var route models.RouteWrapper
	err := c.BindJSON(route)
	if err != nil {

	}

	// TODO Build image & push from code.

	// TODO create Functions's route
	functions.CreateRoute(route)
}
