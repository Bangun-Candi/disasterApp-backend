package controllers

import (
	"net/http"
	"users/models"
	"users/services"

	"github.com/gin-gonic/gin"
)

func GetRescuersCategory(c *gin.Context) {
	var request struct {
		UserEmail string `json:"userEmail"`
		UserName  string `json:"userName"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	rescuers, err := services.GetRescuersCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get rescuers category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": gin.H{"categoryRescuers": rescuers}})
}

func SendRescueDisaster(c *gin.Context) {
	var request models.RescueRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	rescuers, err := services.SendRescueDisaster(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": gin.H{"listRescuers": rescuers}})
}
