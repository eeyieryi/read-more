package main

import (
	"net/http"
	"read-more-backend/api"
	"read-more-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := database.Connect()
	database.Migrate(db)
	router := gin.Default()
	router.Use(cors.Default())
	router.StaticFS("/audios", http.Dir("./_data/audios"))
	api.Setup(router, db)
	router.Run()
}
