package routes

import (
	"net/http"

	"github.com/Chidinma21/Events-Booking-API/models"
	"github.com/Chidinma21/Events-Booking-API/utils"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse signup request data"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse signup request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
