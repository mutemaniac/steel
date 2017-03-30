package main

import (
	"fmt"

	"github.com/mutemaniac/steel/config"
	"github.com/mutemaniac/steel/docker"
	"github.com/mutemaniac/steel/functions"
	"github.com/mutemaniac/steel/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	//Login dockerhub
	if config.DockerHubPwd != "" {
		err := docker.Login(config.DockerHubServer, config.DockerHubUser, config.DockerHubPwd)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.POST("/route", createRoute)
	}

	router.Run(":8081")
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
