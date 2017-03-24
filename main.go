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
	var route models.ExRouteWrapper
	err := c.BindJSON(&route)
	if err != nil {
		c.JSON(304, gin.H{
			"message": err.Error(),
		})
		return
	}
	r, err := functions.CreateRoute(route)
	if err != nil {
		c.JSON(304, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, r)
	}
}
