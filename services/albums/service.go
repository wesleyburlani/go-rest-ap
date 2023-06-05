package albums

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/models"
	"gorm.io/gorm"
)

type AlbumsService struct {
	db     *gorm.DB
	logger *logrus.Logger
	ctx    context.Context
}

func NewAlbumsService(
	db *gorm.DB,
	logger *logrus.Logger,
) *AlbumsService {
	return &AlbumsService{
		db:     db,
		logger: logger,
		ctx:    context.Background(),
	}
}

func (instance *AlbumsService) WithContext(ctx context.Context) IAlbumsService {
	instance.ctx = ctx
	return instance
}

func (instance *AlbumsService) Create(props models.AlbumProps) models.Album {
	album := models.Album{
		Title:  props.Title,
		Artist: props.Artist,
		Price:  props.Price,
	}
	instance.db.WithContext(instance.ctx).Create(&album)

	instance.logger.WithContext(instance.ctx).WithFields(logrus.Fields{
		"id": album.ID,
	}).Debug("create album")

	return album
}

func (instance *AlbumsService) List(page int, limit int) []models.Album {
	albums := []models.Album{}
	instance.
		db.
		WithContext(instance.ctx).
		Model(&models.Album{}).
		Offset(page * limit).
		Limit(limit).
		Find(&albums)

	instance.logger.WithContext(instance.ctx).WithFields(logrus.Fields{
		"resultLength": len(albums),
	}).Debug("get albums")

	return albums
}

func (instance *AlbumsService) Get(id uint) (models.Album, error) {
	album := models.Album{
		ID: id,
	}
	result := instance.db.WithContext(instance.ctx).First(&album)

	instance.logger.WithContext(instance.ctx).WithFields(logrus.Fields{
		"id":    album.ID,
		"error": result.Error,
	}).Debug("get album")

	return album, result.Error
}

func (instance *AlbumsService) Update(id uint, props models.AlbumProps) (models.Album, error) {
	album := models.Album{
		ID:     id,
		Title:  props.Title,
		Artist: props.Artist,
		Price:  props.Price,
	}
	result := instance.db.WithContext(instance.ctx).Save(&album)

	instance.logger.WithContext(instance.ctx).WithFields(logrus.Fields{
		"id":    album.ID,
		"error": result.Error,
	}).Debug("update album")

	return album, result.Error
}
