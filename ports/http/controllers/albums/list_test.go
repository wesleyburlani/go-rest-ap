package albums_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/models"
	http_controller_albums "github.com/wesleyburlani/go-rest-api/ports/http/controllers/albums"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"

	"github.com/gin-gonic/gin"
)

func setupGetAlbumsTest() (*gin.Engine, *service_albums.MockAlbumsService) {
	router := gin.New()
	svc := service_albums.NewMockAlbumsService()
	logger := logrus.New()
	logger.Out = io.Discard
	controller := http_controller_albums.NewListAlbumsController(logger, svc)
	router.Handle(controller.Method(), controller.RelativePath(), controller.Handle)
	return router, svc
}

func listAlbumsRequest(router *gin.Engine, page string, limit string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums?page=%s&limit=%s", page, limit), nil)
	router.ServeHTTP(w, req)
	return w
}

func TestListAlbums_Success(t *testing.T) {
	router, svc := setupGetAlbumsTest()

	svc.Albums = []models.Album{}
	for i := 0; i < 5; i++ {
		svc.Create(models.AlbumProps{
			Title:  "test",
			Artist: "test",
			Price:  1.0,
		})
	}

	response := listAlbumsRequest(router, "0", "10")
	expected := 200
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}

	var result []models.Album
	err := json.Unmarshal(response.Body.Bytes(), &result)

	if err != nil {
		t.Error(err)
	}

	expectLength := 5
	if len(result) != expectLength {
		t.Errorf("Expected %d, received %d", expectLength, len(result))
	}
}

func TestListAlbums_InvalidQueryParamPage(t *testing.T) {
	router, _ := setupGetAlbumsTest()

	response := listAlbumsRequest(router, "-1", "1")
	expected := 400
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}

func TestListAlbums_InvalidQueryParamLimit(t *testing.T) {
	router, _ := setupGetAlbumsTest()

	response := listAlbumsRequest(router, "0", "-1")
	expected := 400
	if response.Code != expected {
		t.Errorf("Expected %d, received %d", expected, response.Code)
	}
}
