package albums_controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wesleyburlani/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func getAlbumRequest(router *gin.Engine, id string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%s", id), nil)
	router.ServeHTTP(w, req)
	return w
}

func TestGetAlbum(t *testing.T) {
	album := mockAlbumsService.CreateAlbum(models.AlbumProps{
		Title:  "test",
		Artist: "test",
		Price:  1.0,
	})

	router := setupRouter()
	response := getAlbumRequest(router, album.ID)
	expected := 200
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}

func TestGetAlbum_NotFound(t *testing.T) {
	router := setupRouter()
	response := getAlbumRequest(router, "not-found-id")
	expected := 404
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}
