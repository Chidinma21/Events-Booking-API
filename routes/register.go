package routes

import (
	"net/http"
	"strconv"

	"github.com/Chidinma21/Events-Booking-API/models"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	registration := models.Registration{
		UserID:  userId,
		EventID: eventId,
	}

	err = registration.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save registration"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event registration created successfully"})
}

func Cancel(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	registration := models.Registration{
		UserID:  userId,
		EventID: eventId,
	}

	err = registration.CancelRegistration()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event registration cancelled successfully"})

}
