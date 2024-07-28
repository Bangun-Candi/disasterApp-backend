package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"users/models"
	"users/services"
	"users/utils"

	"github.com/gin-gonic/gin"
)

func GetCurrentStatus(c *gin.Context) {
	var request struct {
		UserEmail         string `json:"userEmail"`
		UserName          string `json:"userName"`
		LongitudeLocation string `json:"longitudeLocation"`
		LatitudeLocation  string `json:"latitudeLocation"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": "Invalid request"})
		return
	}

	status, err := services.GetCurrentStatus(request.UserEmail, request.LongitudeLocation, request.LatitudeLocation)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"statusCode": "404", "isError": true, "message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": "200",
		"isError":    false,
		"message":    "SUCCESS",
		"data":       status,
	})
}

func SendRealtimeLocation(c *gin.Context) {
	var request struct {
		UserEmail         string `json:"userEmail"`
		UserName          string `json:"userName"`
		LongitudeLocation string `json:"longitudeLocation"`
		LatitudeLocation  string `json:"latitudeLocation"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": "Invalid request"})
		return
	}

	lat, err := strconv.ParseFloat(request.LatitudeLocation, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": "Invalid latitude"})
		return
	}

	long, err := strconv.ParseFloat(request.LongitudeLocation, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": "400", "isError": true, "message": "Invalid longitude"})
		return
	}

	// Get the user ID based on the email
	var user models.User
	db := utils.GetDB()
	err = db.QueryRow("SELECT id FROM users WHERE email = ?", request.UserEmail).Scan(&user.ID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"statusCode": "404", "isError": true, "message": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": "Database error"})
		return
	}

	locationRequest := models.RealTimeLocation{
		UserID:    user.ID,
		Latitude:  lat,
		Longitude: long,
	}

	location, err := services.SendRealtimeLocation(locationRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": "500", "isError": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": "200",
		"isError":    false,
		"message":    "SUCCESS",
		"data": gin.H{
			"userEmail":         request.UserEmail,
			"userName":          request.UserName,
			"longitudeLocation": request.LongitudeLocation,
			"latitudeLocation":  request.LatitudeLocation,
			"locationName":      location.LocationName,
			"disasterCode":      location.DisasterCode,
			"disasterName":      location.DisasterName,
		},
	})
}
