package routes

import (
	"net/http"
	"strconv"

	"github.com/Chidinma21/Events-Booking-API/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	userId := ctx.GetInt64("userId")

	event.UserID = userId

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "created event", "event": event})

}

func UpdateEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventId

	if userId != event.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update this event"})
		return
	}
	
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})

}

func DeleteEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if userId != event.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
