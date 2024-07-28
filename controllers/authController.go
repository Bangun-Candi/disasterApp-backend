package controllers

import (
	"users/models"
	"users/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterOnboarding(c *gin.Context) {
	var request models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	err := services.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS"})
}

func Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	user, err := services.LoginUser(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"isError": true, "message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": "200",
		"isError":    false,
		"message":    "SUCCESS",
		"data": gin.H{
			"userID":        user.ID,
			"userEmail":     user.Email,
			"userName":      user.Name,
			"phoneNumber":   user.PhoneNumber,
		},
	})
}
