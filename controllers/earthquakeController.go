package controllers

import (
	"net/http"
	"users/services"

	"github.com/gin-gonic/gin"
)

func FetchEarthquakeData(c *gin.Context) {
	earthquake, err := services.FetchBMKGEarthquakeData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to fetch earthquake data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": earthquake})
}
