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

func EntryCreateOne(c *gin.Context) {
	funcName := "EntryCreateOne"

	var dto models.CreateEntryDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		BadRequest(c)
		return
	}

	if utils.IsEmpty(dto.Title) || utils.IsEmpty(dto.Transcription) {
		log.Println(utils.GetLogMessage(funcName, 1, "(title) and (transcription) must be provided"))
		BadRequest(c)
		return
	}

	var hasCollectionId, hasNewCollectionTitle bool

	hasCollectionId = !utils.IsEmpty(dto.CollectionID)
	hasNewCollectionTitle = !utils.IsEmpty(dto.CollectionTitle)

	if !hasCollectionId && !hasNewCollectionTitle {
		log.Println(utils.GetLogMessage(funcName, 2, "(collectionId) or *new* (collectionTitle) must be provided"))
		BadRequest(c)
		return
	}

	var parsedId uuid.UUID
	if hasCollectionId {
		var err error
		parsedId, err = uuid.Parse(dto.CollectionID)
		if err != nil || parsedId == uuid.Nil {
			log.Println(utils.GetLogMessage(funcName, 3, err))
			BadRequest(c)
			return
		}
	}

	entry := models.CreateEntryFromDto(dto, parsedId)
	if err := entry.CreateOne(models.Database, dto.CollectionTitle); err != nil {
		log.Println(utils.GetLogMessage(funcName, 4, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func EntryFindOne(c *gin.Context) {
	funcName := "EntryFindOne"

	id := c.MustGet("id").(uuid.UUID)

	var entry models.Entry

	var with models.Entry
	with.ID = id

	if err := entry.FindOne(models.Database, &with); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			NotFound(c)
			return
		}
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func EntryUpdateOne(c *gin.Context) {
	funcName := "EntryUpdateOne"

	id := c.MustGet("id").(uuid.UUID)

	var dto models.Entry

	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		BadRequest(c)
		return
	}

	if dto.CollectionID == uuid.Nil || utils.IsEmpty(dto.Title) || utils.IsEmpty(dto.Transcription) {
		log.Println(utils.GetLogMessage(funcName, 1, "(title) and (transcription) and (collectionId) must be provided"))
		BadRequest(c)
		return
	}

	var entry models.Entry
	entry.ID = id

	if err := entry.UpdateOne(models.Database, &dto); err != nil {
		log.Println(utils.GetLogMessage(funcName, 2, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func EntryDeleteOne(c *gin.Context) {
	funcName := "EntryDeleteOne"

	id := c.MustGet("id").(uuid.UUID)

	var entry models.Entry
	entry.ID = id

	if err := entry.DeleteOne(models.Database); err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}
