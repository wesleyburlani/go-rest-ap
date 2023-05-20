package albums_service

import "test/web-service/models"

type IAlbumsService interface {
	CreateAlbum(props models.AlbumProps) models.Album
	GetAlbum(id string) (models.Album, error)
	GetAlbums(page int, limit int) []models.Album
	UpdateAlbum(id string, props models.AlbumProps) (models.Album, error)
}
