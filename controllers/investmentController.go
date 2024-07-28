package controllers

import (
	"users/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInvestmentReferences(c *gin.Context) {
	investments, err := services.GetInvestmentReferences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get investment references"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": investments})
}
