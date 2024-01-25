package usecases

import (
	"example/web-service-gin/src/core/entities"
	"example/web-service-gin/src/infra/db/repositories"
)

func CreateAlbumUseCase(album entities.Album) entities.Album {
	newAlbum := repositories.CreateAlbum(album)
	return *newAlbum
}
