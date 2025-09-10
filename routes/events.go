package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
)

func getEvents(ctx *gin.Context) {
	// get all events from DB
	events, err := models.GetEvents()

	if err != nil {
		log.Printf("🫴 getEvents() handler function has error when models.GetEvents() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Do not have access to records."})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data": gin.H{
				"events": events,
			},
		})
}

func getEvent(ctx *gin.Context) {
	// get id of event from URL with middleware
	eventId := ctx.GetInt("id_param")

	// get event from DB
	event, err := models.GetEvent(eventId)
	if err != nil {
		log.Printf("🫴 getEvent() handler function has error when models.GetEvent() call: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "The desired record was not found."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"event": event,
		},
	})
}

func createEvents(ctx *gin.Context) {
	var event models.Event

	// get user id from token with middleware
	event.UserId = ctx.GetInt("user_id")

	// get data from body of request
	if err := ctx.ShouldBindJSON(&event); err != nil {
		log.Printf("🫴 createEvents() handler function has error when read body of request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body"})
		return
	}

	// save event in DB
	if err := event.Save(); err != nil {
		// TODO: use these for create clear error
		log.Printf("🫴 createEvents() handler function has error when event.Save() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something is wrong"})
		return
	}

	// ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful", "event": event})
	ctx.JSON(http.StatusCreated, gin.H{"success": true, "message": "create new event successful"})
}

func updateEvent(ctx *gin.Context) {
	var event models.Event
	// get id of event from URL with middleware
	event.ID = ctx.GetInt("id_param")

	// get user id from token with middleware
	event.UserId = ctx.GetInt("user_id")

	// get event data from DB
	eventModel, err := models.GetEvent(event.ID)
	if err != nil {
		log.Printf("🫴 updateEvent() handler function has error when models.GetEvent() call: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "The desired record was not found."})
		return
	}

	// check accessibility
	if eventModel.UserId != event.UserId {
		log.Println("🫴 updateEvent() handler function has error because of due to a mismatch between the user ID and the event owner")
		ctx.JSON(http.StatusForbidden, gin.H{"success": false, "message": "You don't access to update this event"})
		return
	}

	// get data from body of request
	if err := ctx.ShouldBindJSON(&event); err != nil {
		log.Printf("🫴 updateEvent() handler function has error when read body of request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body."})
		return
	}

	// update event in DB
	if err := event.Update(); err != nil {
		// TODO: use these for create clear error
		log.Printf("🫴 updateEvent() handler function has error when event.Update() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something is wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Update event successful"})
}

func deleteEvent(ctx *gin.Context) {
	// get id of event from URL with middleware
	eventId := ctx.GetInt("id_param")

	// get event from DB
	event, err := models.GetEvent(eventId)
	if err != nil {
		log.Printf("🫴 deleteEvent() handler function has error when models.GetEvent() call: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "The desired record was not found."})
		return
	}

	// check accessibility
	if event.UserId != ctx.GetInt("user_id") {
		log.Println("🫴 deleteEvent() handler function has error because of due to a mismatch between the user ID and the event owner")
		ctx.JSON(http.StatusForbidden, gin.H{"success": false, "message": "You don't access to update this event"})
		return
	}

	// delete event from DB
	if err := event.Delete(); err != nil {
		log.Printf("🫴 deleteEvent() handler function has error when event.Delete() call: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Could not delete desired record"})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "delete event successful."})
}
