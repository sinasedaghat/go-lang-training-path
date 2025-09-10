package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
)

func registerEvent(ctx *gin.Context) {
	// get id of event from URL with middleware
	eventId := ctx.GetInt("id_param")

	// get user id from token with middleware
	userId := ctx.GetInt("user_id")

	// get event data from DB
	event, err := models.GetEvent(eventId)
	if err != nil {
		log.Printf("🫴 registerEvent() handler function has error when models.GetEvent() call: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "The desired record was not found."})
		return
	}

	// check user owns the event
	if event.UserId == userId {
		log.Println("🫴 registerEvent() handler function has error because of due to a match between the user ID and the event owner")
		ctx.JSON(http.StatusForbidden, gin.H{"success": false, "message": "You are owner of desired event."})
		return
	}

	// register to event in DB
	if err := event.Register(userId); err != nil {
		// TODO: use these for create clear error
		log.Printf("🫴 registerEvent() handler function has error when event.Register() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Something is wrong."})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "You have registered for the desired event."})
}

func unregisterEvent(ctx *gin.Context) {

	var event models.Event
	// get id of event from URL with middleware
	/*
		you can use this `value, exist := ctx.Get("id_param")`
		for know `id_param` set and use `eventId := value.(int)` for get correct value.
	*/
	event.ID = ctx.GetInt("id_param")

	// get user id from token with middleware
	userId := ctx.GetInt("user_id")

	// unregister to event in DB
	if err := event.Unregister(userId); err != nil {
		// TODO: use these for create clear error
		log.Printf("🫴 unregisterEvent() handler function has error when event.Unregister() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Event registration could not be canceled."})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Event registration canceled."})
}
