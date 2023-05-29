package http_api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/http_api"
	"github.com/wesleyburlani/go-rest-api/models"
	"github.com/wesleyburlani/go-rest-api/services"
)

func setupGetAlbumTest() (*gin.Engine, *services.MockAlbumsService) {
	router := gin.New()
	svc := services.NewMockAlbumsService()
	controller := http_api.NewGetAlbumController(svc)
	router.Handle(controller.Method(), controller.RelativePath(), controller.Handle)
	return router, svc
}

func getAlbumRequest(router *gin.Engine, id string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%s", id), nil)
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

	response := getAlbumRequest(router, "not-found-id")
	expected := 404
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}
