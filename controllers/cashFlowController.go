package controllers

import (
	"net/http"
	"strconv"
	"time"
	"users/services"

	"github.com/gin-gonic/gin"
)

func GetCashFlowReport(c *gin.Context) {
	var request struct {
		UserID    string `json:"userID"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	userID, err := strconv.Atoi(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid User ID"})
		return
	}

	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid Start Date"})
		return
	}
	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid End Date"})
		return
	}

	cashFlows, err := services.GetCashFlowReport(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get cash flow report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": cashFlows})
}
