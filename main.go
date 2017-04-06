package main

import (
	"context"
	"fmt"
	"time"

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
		async := v1.Group("async")
		{
			async.POST("/route", asyncCreateRoute)
		}
	}

	router.Run(":8081")
}
func asyncCreateRoute(c *gin.Context) {
	var route models.AsyncRouteWrapper
	err := c.BindJSON(&route)
	if err != nil {
		c.JSON(304, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx, _ := context.WithCancel(context.Background())
	taskid, err := functions.AsyncCreateRoute(ctx, route)
	if err != nil {
		c.JSON(304, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"taskid": taskid,
		})
	}
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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	r, err := functions.CreateRoute(ctx, route)
	if err != nil {
		c.JSON(304, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, r)
	}

}
