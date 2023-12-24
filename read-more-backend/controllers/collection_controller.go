package controllers

import (
	"errors"
	"log"
	"net/http"
	"read-more-backend/models"
	"read-more-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CollectionCreateOne(c *gin.Context) {
	funcName := "CollectionCreateOne"

	var collection models.Collection

	if err := c.ShouldBindJSON(&collection); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		BadRequest(c)
		return
	}

	if utils.IsEmpty(collection.Title) {
		log.Println(utils.GetLogMessage(funcName, 1, "collection title is empty"))
		BadRequest(c)
		return
	}

	if err := collection.CreateOne(models.Database); err != nil {
		log.Println(utils.GetLogMessage(funcName, 2, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, collection)
}

func CollectionFindAll(c *gin.Context) {
	funcName := "CollectionFindAll"

	collections, err := (&models.Collection{}).FindAll(models.Database)
	if err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, collections)
}

func CollectionFindOne(c *gin.Context) {
	funcName := "CollectionFindOne"

	id := c.MustGet("id").(uuid.UUID)

	var collection models.Collection

	var with models.Collection
	with.ID = id

	if err := collection.FindOne(models.Database, &with); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			NotFound(c)
			return
		}
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, collection)
}

func CollectionDeleteOne(c *gin.Context) {
	funcName := "CollectionDeleteOne"

	id := c.MustGet("id").(uuid.UUID)

	var collection models.Collection
	collection.ID = id

	if err := collection.DeleteOne(models.Database); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, collection)
}
