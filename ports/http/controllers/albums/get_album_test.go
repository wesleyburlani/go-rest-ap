package albums_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/models"
	http_controller_albums "github.com/wesleyburlani/go-rest-api/ports/http/controllers/albums"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
)

func setupGetAlbumTest() (*gin.Engine, *service_albums.MockAlbumsService) {
	router := gin.New()
	svc := service_albums.NewMockAlbumsService()
	controller := http_controller_albums.NewGetAlbumController(svc)
	router.Handle(controller.Method(), controller.RelativePath(), controller.Handle)
	return router, svc
}

func getAlbumRequest(router *gin.Engine, id uint) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%d", id), nil)
	router.ServeHTTP(w, req)
	return w
}

func TestGetAlbum_Success(t *testing.T) {
	router, svc := setupGetAlbumTest()

	album := svc.CreateAlbum(models.AlbumProps{
		Title:  "test",
		Artist: "test",
		Price:  1.0,
	})

	response := getAlbumRequest(router, album.ID)
	expected := 200
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}

func TestGetAlbum_NotFound(t *testing.T) {
	router, _ := setupGetAlbumTest()

	response := getAlbumRequest(router, 99999)
	expected := 404
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}
