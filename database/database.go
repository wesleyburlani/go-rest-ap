package database

import (
	"strconv"

	"github.com/wesleyburlani/go-rest-api/models"

	linq "github.com/ahmetb/go-linq/v3"
)

type Database struct {
}

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func NewDatabase() *Database {
	return &Database{}
}

func (instance *Database) GetAlbums(page int, limit int) []models.Album {
	skip := (page) * (limit)
	var results []models.Album
	linq.From(albums).Skip(skip).Take(limit).ToSlice(&results)
	return results
}

func (instance *Database) GetAlbum(id string) (models.Album, bool) {
	album := linq.From(albums).FirstWithT(func(a models.Album) bool {
		return a.ID == id
	})

	if album == nil {
		return models.Album{}, false
	}

	return album.(models.Album), true
}

func (instance *Database) CreateAlbum(props models.AlbumProps) models.Album {
	album := models.Album{
		ID:     strconv.Itoa((len(albums) + 2)),
		Title:  props.Title,
		Artist: props.Artist,
		Price:  props.Price,
	}

	albums = append(albums, album)
	return album
}

func (instance *Database) UpdateAlbum(id string, props models.AlbumProps) (models.Album, bool) {
	for index, album := range albums {
		if album.ID == id {
			albums[index] = models.Album{
				ID:     id,
				Title:  props.Title,
				Artist: props.Artist,
				Price:  props.Price,
			}
			return albums[index], true
		}
	}

	return models.Album{}, false
}
