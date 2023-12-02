package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"read-more-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func collectionCreateOne(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[Error] collectionCreateOne [0]:", err)
		badRequest(c)
		return
	}

	collection := models.Collection{}
	err = json.Unmarshal(data, &collection)
	if err != nil {
		log.Println("[Error] collectionCreateOne [1]:", err)
		badRequest(c)
		return
	}

	err = collection.CreateOne(db)
	if err != nil {
		log.Println("[Error] collectionCreateOne [2]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusCreated, collection)
}

func collectionFindAll(c *gin.Context) {
	collections, err := (&models.Collection{}).FindAll(db)
	if err != nil {
		internalServerError(c)
		return
	}

	c.JSON(http.StatusOK, collections)
}

func collectionFindOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] collectionFindOne [0]:", err)
		badRequest(c)
		return
	}

	collection := &models.Collection{}
	err = collection.FindOne(db, &models.Collection{ID: id})
	if err != nil {
		log.Println("[Error] collectionFindOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, collection)
}

func collectionDeleteOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] collectionDeleteOne [0]:", err)
		badRequest(c)
		return
	}

	collection := &models.Collection{
		ID: id,
	}
	err = collection.DeleteOne(db)
	if err != nil {
		log.Println("[Error] collectionDeleteOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, collection)
}
