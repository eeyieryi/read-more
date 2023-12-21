package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func badRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
	})
}

func internalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
	})
}
