package repositories

import (
	"errors"
	"example/web-service-gin/src/core/entities"
)

var albums = []entities.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbumById(id string) (*entities.Album, error) {
	for _, album := range albums {
		if album.ID == id {

			return &album, nil
		}
	}

	return nil, errors.New("Album with id " + id + " not found.")
}

func GetAlbums() *[]entities.Album {
	return &albums
}

func CreateAlbum(newAlbum entities.Album) *entities.Album {
	albums = append(albums, newAlbum)
	return &newAlbum
}
