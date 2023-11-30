package api

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("audioFile")
	if err != nil {
		log.Println("[Error] UploadHandler [0]:", err)
		badRequest(c)
		return
	}

	entryTitle := c.PostForm("entryTitle")
	if len(entryTitle) == 0 {
		log.Println("[Error] UploadHandler [1]:", "empty entry title")
		badRequest(c)
		return
	}

	filename := fmt.Sprintf("temp/%s.mp3", entryTitle)

	c.SaveUploadedFile(file, filepath.Join("_data", "audios", filename))

	c.JSON(http.StatusOK, gin.H{
		"audioFilename": filename,
	})
}
