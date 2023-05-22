package albums_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wesleyburlani/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func executeGetAlbumsRequest(router *gin.Engine, page string, limit string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums?page=%s&limit=%s", page, limit), nil)
	router.ServeHTTP(w, req)
	return w
}

func TestGetAlbums(t *testing.T) {
	mockAlbumsService.Albums = []models.Album{}
	for i := 0; i < 5; i++ {
		mockAlbumsService.CreateAlbum(models.AlbumProps{
			Title:  "test",
			Artist: "test",
			Price:  1.0,
		})
	}
	router := setupRouter()
	response := executeGetAlbumsRequest(router, "0", "10")
	expected := 200
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}

	var result []models.Album
	json.Unmarshal(response.Body.Bytes(), &result)
	expectLength := 5
	if len(result) != expectLength {
		t.Errorf("Expected %d, received %d", expectLength, len(result))
	}
}

func TestGetAlbums_InvalidQueryParamPage(t *testing.T) {
	router := setupRouter()
	response := executeGetAlbumsRequest(router, "-1", "1")
	expected := 400
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}

func TestGetAlbums_InvalidQueryParamLimit(t *testing.T) {
	router := setupRouter()
	response := executeGetAlbumsRequest(router, "0", "-1")
	expected := 400
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}
