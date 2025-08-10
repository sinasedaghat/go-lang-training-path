package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/models"
	"sample-url.com/REST-API/utils"
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

func createEvents(ctx *gin.Context) {
	// Authentication action
	token := ctx.Request.Header.Get("Authentication")

	if token == "" {
		fmt.Println("token is empty")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "please sign in."})
		return
	}

	ok, id := utils.ParseTokenAndGetID(token)
	fmt.Println("ok from ParseTokenAndGetID", ok)
	fmt.Println("id from ParseTokenAndGetID", id)
	if !ok {
		fmt.Println("parse ParseTokenAndGetID has error")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "please sign in."})
		return
	}

	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("error from create event read body of request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}

	event.UserId = 1 // TODO: after create user table
	err = event.Save()
	if err != nil {
		fmt.Println("error from create event save function", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}
	// ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful", "event": event})
	ctx.JSON(http.StatusCreated, gin.H{"message": "create new event successful"})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := getId(ctx)

	if err != nil {
		fmt.Println("error from read id from GetEvent function", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid ID parameter"})
		return
	}

	_, err = models.GetEvent(eventId)

	if err != nil {
		fmt.Println("error from get specific event in updateEvent function", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired record was not found."})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println("error from updateEvent read body of request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}
	event.ID = eventId
	event.UserId = 1 // TODO: after create user table
	err = event.Update()
	if err != nil {
		fmt.Println("error from update event in update function", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "update event successful"})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := getId(ctx)

	if err != nil {
		fmt.Println("error from read id from deleteEvent function", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid ID parameter"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		fmt.Println("error from get specific event in deleteEvent function", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "The desired record was not found."})
		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println("error from delete specific event in deleteEvent function", err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not delete desired record"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete event successful."})
}
