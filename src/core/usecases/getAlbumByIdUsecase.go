package usecases

import (
	"example/web-service-gin/src/core/entities"
	"example/web-service-gin/src/infra/db/repositories"
)

func GetAlbumByIdUseCase(id string) (*entities.Album, error) {
	album, err := repositories.GetAlbumById(id)

	if err != nil {
		return nil, err
	}

	return album, nil
}
