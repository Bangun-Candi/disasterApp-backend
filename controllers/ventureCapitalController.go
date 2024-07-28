package controllers

import (
	"users/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVentureCapital(c *gin.Context) {
	vcs, err := services.GetVentureCapital()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isError": true, "message": "Failed to get venture capital list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": "200", "isError": false, "message": "SUCCESS", "data": vcs})
}
