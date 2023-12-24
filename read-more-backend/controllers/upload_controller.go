package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"read-more-backend/utils"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	funcName := "UploadHandler"

	file, err := c.FormFile("audioFile")
	if err != nil {
		log.Println(utils.GetLogMessage(funcName, 0, err))
		BadRequest(c)
		return
	}

	entryTitle := c.PostForm("entryTitle")
	if utils.IsEmpty(entryTitle) {
		log.Println(utils.GetLogMessage(funcName, 1, "empty entry title"))
		BadRequest(c)
		return
	}

	filename := fmt.Sprintf("temp/%s.mp3", entryTitle)

	if err := c.SaveUploadedFile(file, filepath.Join("_data", "audios", filename)); err != nil {
		log.Println(utils.GetLogMessage(funcName, 2, err))
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"audioFilename": filename,
	})
}
