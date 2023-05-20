package database_test

import (
	"test/web-service/database"
	"test/web-service/models"
	"testing"
)

func TestGetAlbum(t *testing.T) {
	instance := database.NewDatabase()
	createdAlbum := instance.CreateAlbum(models.AlbumProps{
		Title:  "test",
		Artist: "test",
		Price:  1.0,
	})
	expected := true
	_, exists := instance.GetAlbum(createdAlbum.ID)
	if exists != expected {
		t.Errorf("Expected %t, received %t", expected, exists)
	}
}

func TestGetAlbum_NotFoundElement(t *testing.T) {
	instance := database.NewDatabase()
	expected := false
	_, exists := instance.GetAlbum("not-found-id")
	if exists != expected {
		t.Errorf("Expected %t, received %t", expected, exists)
	}
}

func TestUpdateAlbum_NotFoundElement(t *testing.T) {
	instance := database.NewDatabase()
	_, updated := instance.UpdateAlbum("not-found-id", models.AlbumProps{})
	if updated == true {
		t.Errorf("Expected false, received true")
	}
}

func TestUpdateAlbum(t *testing.T) {
	instance := database.NewDatabase()

	createdAlbum := instance.CreateAlbum(models.AlbumProps{
		Title:  "test",
		Artist: "test",
		Price:  1.0,
	})

	album, updated := instance.UpdateAlbum(createdAlbum.ID, models.AlbumProps{
		Title:  "test1",
		Artist: "test1",
		Price:  2.0,
	})
	if updated != true {
		t.Errorf("Expected true, received %v", updated)
	}

	if album.Title != "test1" {
		t.Errorf("Expected test1, received %s", album.Title)
	}

	v, _ := instance.GetAlbum(createdAlbum.ID)

	if v.Title != "test1" {
		t.Errorf("Expected test1, received %s", v.Title)
	}
}
