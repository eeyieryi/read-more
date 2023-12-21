package controllers

import (
	"log"
	"net/http"
	"read-more-backend/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CollectionCreateOne(c *gin.Context) {
	var err error
	var collection models.Collection

	err = c.ShouldBindJSON(&collection)
	if err != nil {
		log.Println("[Error] collectionCreateOne [0]:", err)
		badRequest(c)
		return
	}

	collection.Title = strings.TrimSpace(collection.Title)
	if isEmpty(collection.Title) {
		log.Println("[Error] collectionCreateOne [1]")
		badRequest(c)
		return
	}

	err = collection.CreateOne(models.Database)
	if err != nil {
		log.Println("[Error] collectionCreateOne [2]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusCreated, collection)
}

func CollectionFindAll(c *gin.Context) {
	collections, err := (&models.Collection{}).FindAll(models.Database)
	if err != nil {
		internalServerError(c)
		return
	}

	c.JSON(http.StatusOK, collections)
}

func CollectionFindOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] collectionFindOne [0]:", err)
		badRequest(c)
		return
	}

	collection := &models.Collection{}
	err = collection.FindOne(models.Database, &models.Collection{ID: id})
	if err != nil {
		log.Println("[Error] collectionFindOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, collection)
}

func CollectionDeleteOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] collectionDeleteOne [0]:", err)
		badRequest(c)
		return
	}

	collection := &models.Collection{
		ID: id,
	}
	err = collection.DeleteOne(models.Database)
	if err != nil {
		log.Println("[Error] collectionDeleteOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, collection)
}
