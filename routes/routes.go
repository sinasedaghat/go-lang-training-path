package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
	})

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// server.GET("/users", getUsers) // TODO: we need high level role (admin) for get all users and change some thing in user data
	server.POST("/sign-up", createUsers)
	server.POST("/sign-in", signInUsers)
}
