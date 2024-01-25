package routes

import (
	"example/web-service-gin/src/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.POST("/albums", controllers.PostAlbum)
	return router
}
