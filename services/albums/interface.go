package albums_service

import "github.com/wesleyburlani/go-rest-api/models"

type IAlbumsService interface {
	CreateAlbum(props models.AlbumProps) models.Album
	GetAlbum(id string) (models.Album, error)
	GetAlbums(page int, limit int) []models.Album
	UpdateAlbum(id string, props models.AlbumProps) (models.Album, error)
}
