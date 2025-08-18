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
		fmt.Println("🫴 error from give events in the getEvents handler:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId := ctx.GetInt("id_param")
	event, err := models.GetEvent(eventId)

	if err != nil {
		fmt.Println("🫴 error from give event in the getEvent handler:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired record was not found."})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvents(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")

	var event models.Event
	event.UserId = userId

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("🫴 error from read body of request in the createEvents handler:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}

	err = event.Save()
	if err != nil {
		fmt.Println("🫴 error from event.save() in the createEvents handler:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}

	// ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful", "event": event})
	ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful"})
}

func updateEvent(ctx *gin.Context) {
	eventId := ctx.GetInt("id_param")
	eventModel, err := models.GetEvent(eventId)

	if err != nil {
		fmt.Println("🫴 error from give event in the updateEvent handler:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired record was not found."})
		return
	}

	if eventModel.UserId != ctx.GetInt("user_id") {
		fmt.Println("🫴 error from check owner of event in the updateEvent handler")
		ctx.JSON(http.StatusForbidden, gin.H{"message": "you don't access to update this event"})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("🫴 error from read body of request in the updateEvent handler:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}
	event.ID = eventId
	event.UserId = ctx.GetInt("user_id")
	err = event.Update()
	if err != nil {
		fmt.Println("🫴 error from event.update() in the updateEvent handler:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "update event successful"})
}

func deleteEvent(ctx *gin.Context) {
	eventId := ctx.GetInt("id_param")
	event, err := models.GetEvent(eventId)

	if err != nil {
		fmt.Println("🫴 error from give event in the deleteEvent handler:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired record was not found."})
		return
	}

	if event.UserId != ctx.GetInt("user_id") {
		fmt.Println("🫴 error from check owner of event in the deleteEvent handler")
		ctx.JSON(http.StatusForbidden, gin.H{"message": "you don't access to update this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println("🫴 error from event.delete() in the deleteEvent handler:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not delete desired record"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete event successful."})
}
