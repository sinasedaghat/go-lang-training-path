package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
	})
	// HINT: use middleware for one request
	// server.POST("/events", middleware.Authentication, createEvents)

	// HINT: one approach for create group of requests
	// auth := server.Group("/")
	// auth.Use(middleware.Authentication)
	// auth.POST("/events", createEvents)

	// TODO: after create many middleware create nested groups for router
	auth := server.Group("/", middleware.Authentication)
	{
		auth.POST("/events", createEvents)
		auth.PUT("/events/:id", updateEvent)    // TODO: create new middleware for give `id` from url
		auth.DELETE("/events/:id", deleteEvent) // TODO: create new middleware for give `id` from url
	}

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// server.GET("/users", getUsers) // TODO: we need high level role (admin) for get all users and change some thing in user data
	server.POST("/sign-up", createUsers)
	server.POST("/sign-in", signInUsers)
}
