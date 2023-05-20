package albums_service

import (
	"errors"
	"fmt"
	"test/web-service/database"
	"test/web-service/models"
)

type AlbumsService struct {
	database *database.Database
}

func NewAlbumsService(
	database *database.Database,
) *AlbumsService {
	return &AlbumsService{
		database,
	}
}

func (instance *AlbumsService) CreateAlbum(props models.AlbumProps) models.Album {
	return instance.database.CreateAlbum(props)
}

func (instance *AlbumsService) GetAlbums(page int, limit int) []models.Album {
	return instance.database.GetAlbums(page, limit)
}

func (instance *AlbumsService) GetAlbum(id string) (models.Album, error) {
	album, exists := instance.database.GetAlbum(id)
	if exists == false {
		return models.Album{}, errors.New(fmt.Sprintf("album with id %s not found", id))
	}
	return album, nil
}

func (instance *AlbumsService) UpdateAlbum(id string, props models.AlbumProps) (models.Album, error) {
	album, updated := instance.database.UpdateAlbum(id, props)
	if updated == false {
		return models.Album{}, errors.New(fmt.Sprintf("album with id %s not found", id))
	}
	return album, nil
}
