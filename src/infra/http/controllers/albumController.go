package controllers

import (
	"example/web-service-gin/src/core/entities"
	"example/web-service-gin/src/core/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums := usecases.GetAlbumsUseCase()
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var album entities.Album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	newAlbum := usecases.CreateAlbumUseCase(album)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := usecases.GetAlbumByIdUseCase(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, album)

}
