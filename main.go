package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
)

func main() {
	server := gin.Default()
	defer server.Run(":8080") // Don't call it without "defer" because none of the endpoints are registered. Or call it at the end of the function.

	server.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.GetEvents())
}

func createEvents(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("error from create event read body of request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}

	event.ID = 1
	event.UserId = 1
	event.CreateDate = time.Now()
	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful", "event": event})
}
