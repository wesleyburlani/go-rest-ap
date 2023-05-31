package albums

import "github.com/wesleyburlani/go-rest-api/models"

type IAlbumsService interface {
	CreateAlbum(props models.AlbumProps) models.Album
	GetAlbum(id uint) (models.Album, error)
	GetAlbums(page int, limit int) []models.Album
	UpdateAlbum(id uint, props models.AlbumProps) (models.Album, error)
}
