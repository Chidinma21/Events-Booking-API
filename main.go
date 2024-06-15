package main

import (
	"github.com/Chidinma21/Events-Booking-API/db"
	"github.com/Chidinma21/Events-Booking-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
