package main

import (
	"net/http"
	"read-more-backend/api"
	"read-more-backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDatabase()
	router := gin.Default()
	router.Use(cors.Default())
	router.StaticFS("/audios", http.Dir("./_data/audios"))
	api.Setup(router)
	router.Run()
}
