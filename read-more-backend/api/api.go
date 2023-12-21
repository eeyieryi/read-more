package api

import (
	"read-more-backend/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	api := router.Group("/api")

	v1 := api.Group("/v1")
	currentVersion := v1

	currentVersion.POST("/upload", controllers.UploadHandler)

	collections := currentVersion.Group("collections")
	{
		collections.GET("", controllers.CollectionFindAll)
		collections.POST("", controllers.CollectionCreateOne)

		collections.GET("/:id", controllers.CollectionFindOne)
		collections.DELETE("/:id", controllers.CollectionDeleteOne)
	}

	entries := currentVersion.Group("entries")
	{
		entries.POST("", controllers.EntryCreateOne)

		entries.GET("/:id", controllers.EntryFindOne)
		entries.DELETE("/:id", controllers.EntryDeleteOne)
		entries.PUT("/:id", controllers.EntryUpdateOne)
	}
}
