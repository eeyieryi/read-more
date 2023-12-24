package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
	})
}

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
	})
}
