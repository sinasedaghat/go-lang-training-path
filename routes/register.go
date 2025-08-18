package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
)

func register(ctx *gin.Context) {
	eventId := ctx.GetInt("id_param")
	userId := ctx.GetInt("user_id")

	event, err := models.GetEvent(eventId)

	if err != nil {
		fmt.Println("🫴 error from give event in the register handler:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired event was not found."})
		return
	}

	if event.UserId == userId {
		fmt.Println("🫴 error from check owner of event in the register handler:", err)
		ctx.JSON(http.StatusConflict, gin.H{"message": "You are owner of desired event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Println("🫴 error from event.register (write to database) in the register handler:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something is wrong."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "You have registered for the desired event."})
}

func unregister(ctx *gin.Context) {
	eventId := ctx.GetInt("id_param") // you can use this `value, exist := ctx.Get("id_param")` for know `id_param` set and use `eventId := value.(int)` for get correct value
	userId := ctx.GetInt("user_id")

	var event models.Event
	event.ID = eventId

	err := event.Unregister(userId)

	if err != nil {
		fmt.Println("🫴 error from event.unregister (delete from database) in the unregister handler:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Event registration could not be canceled."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event registration canceled."})
}
