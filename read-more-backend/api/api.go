package api

import (
	"read-more-backend/api/middlewares"
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

		collections.GET("/:id", middlewares.ParseIdFromParams, controllers.CollectionFindOne)
		collections.DELETE("/:id", middlewares.ParseIdFromParams, controllers.CollectionDeleteOne)
	}

	entries := currentVersion.Group("entries")
	{
		entries.POST("", controllers.EntryCreateOne)

		entries.GET("/:id", middlewares.ParseIdFromParams, controllers.EntryFindOne)
		entries.DELETE("/:id", middlewares.ParseIdFromParams, controllers.EntryDeleteOne)
		entries.PUT("/:id", middlewares.ParseIdFromParams, controllers.EntryUpdateOne)
	}
}
