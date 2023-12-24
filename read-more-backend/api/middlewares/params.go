package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ParseIdFromParams(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error while parsing id (%v): (%v)", id, err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if parsedId == uuid.Nil {
		log.Printf("parsed uuid is zero value (%v)", parsedId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Set("id", parsedId)
	c.Next()
}
