package albums

import (
	"context"

	"github.com/wesleyburlani/go-rest-api/models"
	"gorm.io/gorm"
)

type AlbumsService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewAlbumsService(
	db *gorm.DB,
) *AlbumsService {
	return &AlbumsService{
		db:  db,
		ctx: context.Background(),
	}
}

func (instance *AlbumsService) WithContext(ctx context.Context) IAlbumsService {
	instance.ctx = ctx
	return instance
}

func (instance *AlbumsService) CreateAlbum(props models.AlbumProps) models.Album {
	album := models.Album{
		Title:  props.Title,
		Artist: props.Artist,
		Price:  props.Price,
	}
	instance.db.Create(&album)
	return album
}

func (instance *AlbumsService) GetAlbums(page int, limit int) []models.Album {
	albums := []models.Album{}
	instance.db.WithContext(instance.ctx).Model(&models.Album{}).Offset(page * limit).Limit(limit).Find(&albums)
	return albums
}

func (instance *AlbumsService) GetAlbum(id uint) (models.Album, error) {
	album := models.Album{
		ID: id,
	}
	result := instance.db.First(&album)
	return album, result.Error
}

func (instance *AlbumsService) UpdateAlbum(id uint, props models.AlbumProps) (models.Album, error) {
	album := models.Album{
		ID:     id,
		Title:  props.Title,
		Artist: props.Artist,
		Price:  props.Price,
	}
	result := instance.db.Save(&album)
	return album, result.Error
}
