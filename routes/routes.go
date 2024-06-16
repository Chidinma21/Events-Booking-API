package routes

import (
	"github.com/Chidinma21/Events-Booking-API/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	// authenticated.POST("/events/:id/register", Register)
	// authenticated.DELETE("/events/:id/register", Cancel)

	server.POST("/signup", Signup)
	server.POST("/login", Login)
}
