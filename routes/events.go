package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
)

func getId(ctx *gin.Context) (int, error) {
	// return strconv.ParseInt(ctx.Param("id"), 10, 0)
	return strconv.Atoi(ctx.Param("id"))
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		fmt.Println("error from get event getEvents function", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
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
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful", "event": event})
}

func getEvent(ctx *gin.Context) {
	eventId, err := getId(ctx)

	if err != nil {
		fmt.Println("error from read id from GetEvent function", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid ID parameter"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		fmt.Println("error from get specific event in GetEvent function", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired record was not found."})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
