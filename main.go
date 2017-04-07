package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mutemaniac/steel/config"
	"github.com/mutemaniac/steel/docker"
	"github.com/mutemaniac/steel/functions"
	"github.com/mutemaniac/steel/models"
	"github.com/mutemaniac/steel/mqs"
	"gopkg.in/gin-gonic/gin.v1"
)

var MQ *mqs.MemoryMQ

func main() {
	//Login dockerhub
	if config.DockerHubPwd != "" {
		err := docker.Login(config.DockerHubServer, config.DockerHubUser, config.DockerHubPwd)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//Init shoddy mq
	MQ = mqs.NewMemoryMQ()

	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.POST("/route", createRoute)
		async := v1.Group("async")
		{
			async.POST("/route", asyncCreateRoute)
			async.POST("/route/:buildid/cancel", asyncCreateRoute)
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

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	task := mqs.NewSteelTask(route, functions.AsyncCreateRoute)
	ch := make(chan error, 1)
	go func() {
		ch <- MQ.Push(ctx, &task)
	}()
	select {
	case <-ctx.Done():
		c.JSON(304, gin.H{
			"message": "The task queue is full. Please try again later.",
		})
	case err = <-ch:
		if err != nil {
			c.JSON(304, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"taskid": task.Id,
			})
		}
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()
	var r models.ExRouteWrapper
	ch := make(chan error, 1)
	go func() {
		r, err = functions.CreateRoute(ctx, route)
		ch <- err
	}()
	select {
	case <-ctx.Done():
		c.JSON(304, gin.H{
			"message": "Code build timeout",
		})
	case err = <-ch:
		if err != nil {
			c.JSON(304, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, r)
		}
	}
}
