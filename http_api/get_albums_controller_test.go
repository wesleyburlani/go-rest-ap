package http_api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wesleyburlani/go-rest-api/http_api"
	"github.com/wesleyburlani/go-rest-api/models"
	"github.com/wesleyburlani/go-rest-api/services"

	"github.com/gin-gonic/gin"
)

func setupGetAlbumsTest() (*gin.Engine, *services.MockAlbumsService) {
	router := gin.New()
	svc := services.NewMockAlbumsService()
	controller := http_api.NewGetAlbumsController(svc)
	router.Handle(controller.Method(), controller.RelativePath(), controller.Handle)
	return router, svc
}

func getAlbumsRequest(router *gin.Engine, page string, limit string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums?page=%s&limit=%s", page, limit), nil)
	router.ServeHTTP(w, req)
	return w
}

func TestGetAlbums_Success(t *testing.T) {
	router, svc := setupGetAlbumsTest()

	svc.Albums = []models.Album{}
	for i := 0; i < 5; i++ {
		svc.CreateAlbum(models.AlbumProps{
			Title:  "test",
			Artist: "test",
			Price:  1.0,
		})
	}

	response := getAlbumsRequest(router, "0", "10")
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
	router, _ := setupGetAlbumsTest()

	response := getAlbumsRequest(router, "-1", "1")
	expected := 400
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}

func TestGetAlbums_InvalidQueryParamLimit(t *testing.T) {
	router, _ := setupGetAlbumsTest()

	response := getAlbumsRequest(router, "0", "-1")
	expected := 400
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}
