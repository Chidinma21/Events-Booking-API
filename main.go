package main

import (
	"fmt"
	"net/http"

	"github.com/Chidinma21/Events-Booking-API/db"
	"github.com/Chidinma21/Events-Booking-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
	}
	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "created event", "event": event})

}