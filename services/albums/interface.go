package albums

import (
	"context"

	"github.com/wesleyburlani/go-rest-api/models"
)

type IAlbumsService interface {
	WithContext(ctx context.Context) IAlbumsService
	CreateAlbum(props models.AlbumProps) models.Album
	GetAlbum(id uint) (models.Album, error)
	GetAlbums(page int, limit int) []models.Album
	UpdateAlbum(id uint, props models.AlbumProps) (models.Album, error)
}
