package usecases

import (
	"example/web-service-gin/src/core/entities"
	"example/web-service-gin/src/infra/db/repositories"
)

func GetAlbumsUseCase() *[]entities.Album {
	albums := repositories.GetAlbums()
	return albums
}
