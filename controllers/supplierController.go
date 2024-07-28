package controllers

import (
	"net/http"
	"users/services"

	"github.com/gin-gonic/gin"
)

func GetSuppliers(c *gin.Context) {
	suppliers, err := services.GetSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get suppliers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": suppliers})
}
