package controllers

import (
	"net/http"
	"users/services"

	"github.com/gin-gonic/gin"
)

func GenerateQRCode(c *gin.Context) {
	var request struct {
		UserID int   `json:"userID"`
		Amount int64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	qrisTransaction, err := services.GenerateQRCode(request.UserID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to generate QR code", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": qrisTransaction})
}

func ConfirmPayment(c *gin.Context) {
	var request struct {
		QRCode string `json:"qrCode"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isError": true, "message": "Invalid request"})
		return
	}

	qrisTransaction, err := services.ConfirmPayment(request.QRCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to confirm payment", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": qrisTransaction})
}
