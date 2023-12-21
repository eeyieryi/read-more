package controllers

import (
	"log"
	"net/http"
	"read-more-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EntryCreateOne(c *gin.Context) {
	var err error
	var dto models.CreateEntryDto

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		log.Println("[Error] entryCreateOne [0]:", err)
		badRequest(c)
	}

	if len(dto.Title) < 1 || len(dto.Transcription) < 1 {
		log.Println("[Error] entryCreateOne [1])")
		badRequest(c)
	}

	if len(dto.CollectionTitle) < 1 && dto.CollectionID == uuid.Nil {
		log.Println("[Error] entryCreateOne [2]")
		badRequest(c)
	}

	entry := models.Entry{
		Title:         dto.Title,
		Description:   dto.Description,
		URL:           dto.URL,
		Transcription: dto.Transcription,
		AudioFilename: dto.AudioFilename,
		Seen:          false,
		CollectionID:  dto.CollectionID,
	}
	err = entry.CreateOne(models.Database, dto.CollectionTitle)
	if err != nil {
		log.Println("[Error] entryCreateOne [3]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func EntryFindOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] entryFindOne [0]:", err)
		badRequest(c)
		return
	}

	entry := &models.Entry{}
	err = entry.FindOne(models.Database, &models.Entry{ID: id})
	if err != nil {
		log.Println("[Error] entryFindOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func EntryUpdateOne(c *gin.Context) {
	var err error
	var id uuid.UUID

	id, err = uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] entryUpdateOne [0]:", err)
		badRequest(c)
		return
	}

	var dto models.UpdateEntryDto

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		log.Println("[Error] entryUpdateOne [1]:", err)
		badRequest(c)
		return
	}

	if dto.ID == uuid.Nil || id == uuid.Nil {
		log.Println("[Error] entryUpdateOne [2]")
		badRequest(c)
		return
	}

	if dto.ID != id {
		log.Println("[Error] entryUpdateOne [3]")
		badRequest(c)
		return
	}

	entry := models.Entry{
		ID: id,
	}
	err = entry.UpdateOne(models.Database, &dto)
	if err != nil {
		log.Println("[Error] entryUpdateOne [4]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func EntryDeleteOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] entryDeleteOne [0]:", err)
		badRequest(c)
		return
	}

	entry := &models.Entry{
		ID: id,
	}
	err = entry.DeleteOne(models.Database)
	if err != nil {
		log.Println("[Error] entryDeleteOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}
