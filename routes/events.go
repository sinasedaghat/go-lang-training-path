package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		fmt.Println("error from get event getEvents function", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
	}

	ctx.JSON(http.StatusOK, events)
}

func createEvents(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("error from create event read body of request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}

	// event.ID = 1
	event.UserId = 1
	// event.CreateDate = time.Now()
	err = event.Save()
	if err != nil {
		fmt.Println("error from create event save function", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful", "event": event})
}
