package websocket

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/utils"
	"nhooyr.io/websocket"

	"github.com/go-playground/validator/v10"
)

type Controller interface {
	RelativePath() string
	Handle(ctx context.Context, conn *websocket.Conn)
}

type Middleware interface {
	Handle(c *gin.Context)
}

type SocketHandlerFunc func(ctx context.Context, conn *websocket.Conn)

type SocketAuthMessage struct {
	Token string `json:"tokaen"`
}

func SocketHandler(logger *logrus.Logger, handlerFunc SocketHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		connection, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})

		if err != nil {
			logger.WithContext(c.Request.Context()).WithFields(logrus.Fields{
				"error": err.Error(),
			}).Debug("connection error")
			return
		}

		_, raw, err := connection.Read(c.Request.Context())

		if err != nil {
			connection.Close(websocket.StatusPolicyViolation, err.Error())
		}

		var message SocketAuthMessage
		json.Unmarshal(raw, &message)

		err = validator.New().Struct(&message)

		if err != nil {
			logger.WithContext(c.Request.Context()).WithFields(logrus.Fields{"error": err}).Debug("Unauthorized")
			connection.Close(websocket.StatusPolicyViolation, "Unauthorized")
		}

		if message.Token != "123" {
			logger.WithContext(c.Request.Context()).Debug("Unauthorized")
			connection.Close(websocket.StatusPolicyViolation, "Unauthorized")
		}

		handlerFunc(c.Request.Context(), connection)
	}
}

func NewServer(
	middlewares []Middleware,
	controllers []Controller,
	logger *logrus.Logger,
	cfg *utils.Config,
) *gin.Engine {

	if cfg.Mode == utils.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	for _, middleware := range middlewares {
		router.Use(middleware.Handle)
	}

	basePath := "/api/v1"

	v1 := router.Group(basePath)
	for _, controller := range controllers {
		v1.Handle("GET", controller.RelativePath(), SocketHandler(logger, controller.Handle))
	}

	return router
}
