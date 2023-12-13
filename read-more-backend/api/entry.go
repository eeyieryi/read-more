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

func entryCreateOne(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[Error] entryCreateOne [0]:", err)
		badRequest(c)
		return
	}

	createEntryDto := models.CreateEntryDto{}
	err = json.Unmarshal(data, &createEntryDto)
	if err != nil {
		log.Println("[Error] entryCreateOne [1]:", err)
		badRequest(c)
		return
	}

	entry := models.Entry{
		Title:         createEntryDto.Title,
		Description:   createEntryDto.Description,
		URL:           createEntryDto.URL,
		Transcription: createEntryDto.Transcription,
		AudioFilename: createEntryDto.AudioFilename,
		Seen:          false,
		CollectionID:  createEntryDto.CollectionID,
	}
	err = entry.CreateOne(db, createEntryDto.CollectionTitle)
	if err != nil {
		log.Println("[Error] entryCreateOne [2]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func entryFindOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] entryFindOne [0]:", err)
		badRequest(c)
		return
	}

	entry := &models.Entry{}
	err = entry.FindOne(db, &models.Entry{ID: id})
	if err != nil {
		log.Println("[Error] entryFindOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func entryUpdateOne(c *gin.Context) {
	// TODO: Use id from params?
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[Error] entryUpdateOne [0]:", err)
		badRequest(c)
		return
	}

	updateEntryDto := models.UpdateEntryDto{}
	err = json.Unmarshal(data, &updateEntryDto)
	if err != nil {
		log.Println("[Error] entryUpdateOne [1]:", err)
		badRequest(c)
		return
	}

	entry := models.Entry{
		ID: updateEntryDto.ID,
	}

	err = entry.UpdateOne(db, &updateEntryDto)
	if err != nil {
		log.Println("[Error] entryUpdateOne [2]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func entryDeleteOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("[Error] entryDeleteOne [0]:", err)
		badRequest(c)
		return
	}

	entry := &models.Entry{
		ID: id,
	}
	err = entry.DeleteOne(db)
	if err != nil {
		log.Println("[Error] entryDeleteOne [1]:", err)
		badRequest(c)
		return
	}

	c.JSON(http.StatusOK, entry)
}
