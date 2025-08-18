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
		auth.PUT("/events/:id", middleware.URLParameter(), updateEvent)
		auth.DELETE("/events/:id", middleware.URLParameter(), deleteEvent)
		auth.POST("/events/:id/register", middleware.URLParameter(), register)
		auth.DELETE("/events/:id/register", middleware.URLParameter(), unregister)
	}

	server.GET("/events", getEvents)
	server.GET("/events/:id", middleware.URLParameter(), getEvent)

	// server.GET("/users", getUsers) // TODO: we need high level role (admin) for get all users and change some thing in user data
	server.POST("/sign-up", createUsers)
	server.POST("/sign-in", signInUsers)
}
