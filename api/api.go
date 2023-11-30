package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Setup(router *gin.Engine, database *gorm.DB) {
	db = database

	api := router.Group("/api")

	v1 := api.Group("/v1")
	currentVersion := v1

	currentVersion.POST("/upload", uploadHandler)

	collections := currentVersion.Group("collections")
	{
		collections.GET("", collectionFindAll)
		collections.POST("", collectionCreateOne)

		collections.GET("/:id", collectionFindOne)
		collections.DELETE("/:id", collectionDeleteOne)
	}

	entries := currentVersion.Group("entries")
	{
		entries.POST("", entryCreateOne)

		entries.GET("/:id", entryFindOne)
		entries.DELETE("/:id", entryDeleteOne)
		entries.PUT("/:id", entryUpdateOne)
	}
}
