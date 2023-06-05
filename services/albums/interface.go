package albums

import (
	"context"

	"github.com/wesleyburlani/go-rest-api/models"
)

type IAlbumsService interface {
	WithContext(ctx context.Context) IAlbumsService
	Create(props models.AlbumProps) models.Album
	Get(id uint) (models.Album, error)
	List(page int, limit int) []models.Album
	Update(id uint, props models.AlbumProps) (models.Album, error)
}
