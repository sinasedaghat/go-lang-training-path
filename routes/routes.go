package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/constant"
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
		auth.POST("/events/:id/register", middleware.URLParameter(), registerEvent)
		auth.DELETE("/events/:id/register", middleware.URLParameter(), unregisterEvent)
		// auth.GET("/users", middleware., getUsers) // TODO: this method return all users
	}

	server.GET("/events", getEvents)
	server.GET("/events/:id", middleware.URLParameter(), getEvent)

	server.POST("/admin/sign-up", func(ctx *gin.Context) {
		for _, role := range constant.Roles {
			if strings.ToLower(role.Label) == "admin" {
				ctx.Set("role_id", role.Value)
				break
			}
		}
		ctx.Next()
	}, createUser)
	server.POST("/sign-up", createUser)
	server.POST("/sign-in", signInUser)
}
