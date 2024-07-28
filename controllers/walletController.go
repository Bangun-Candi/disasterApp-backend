package controllers

import (
	"users/services"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBalance(c *gin.Context) {
	var request struct {
		UserID    string `json:"userID"`
		UserEmail string `json:"userEmail"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	wallet, err := services.GetBalance(request.UserID, request.UserEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get balance"})
		return
	}

	response := map[string]interface{}{
		"accountFullName": "Muhammad Ramadhani Alfarizi",
		"balance":         wallet.Balance,
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": response})
}

func GetBalanceHistory(c *gin.Context) {
	var request struct {
		UserID    string `json:"userID"`
		UserEmail string `json:"userEmail"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	startDate, _ := time.Parse("2006-01-02", request.StartDate)
	endDate, _ := time.Parse("2006-01-02", request.EndDate)

	history, err := services.GetBalanceHistory(request.UserID, request.UserEmail, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get balance history"})
		return
	}

	response := map[string]interface{}{
		"startDate":          request.StartDate,
		"endDate":            request.EndDate,
		"historyBalanceList": history,
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": response})
}

func GetCompanyGrowth(c *gin.Context) {
	var request struct {
		UserID    string `json:"userID"`
		UserEmail string `json:"userEmail"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	startDate, _ := time.Parse("2006-01-02", request.StartDate)
	endDate, _ := time.Parse("2006-01-02", request.EndDate)

	growth, err := services.GetCompanyGrowth(request.UserID, request.UserEmail, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get company growth"})
		return
	}

	response := map[string]interface{}{
		"startDate":          request.StartDate,
		"endDate":            request.EndDate,
		"balanceInOneMonth":  "8000000000",
		"totalCredit":        "5000000",
		"totalDebit":         "700000000",
		"remainingBalance":   "50000000",
		"historicalBalance":  growth,
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": response})
}

func GetSalesGrowth(c *gin.Context) {
	var request struct {
		UserID    string `json:"userID"`
		UserEmail string `json:"userEmail"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	startDate, _ := time.Parse("2006-01-02", request.StartDate)
	endDate, _ := time.Parse("2006-01-02", request.EndDate)

	sales, err := services.GetSalesGrowth(request.UserID, request.UserEmail, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get sales growth"})
		return
	}

	response := map[string]interface{}{
		"startDate":          request.StartDate,
		"endDate":            request.EndDate,
		"totalTransaction":   "40",
		"historicalBalance":  sales,
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": response})
}
