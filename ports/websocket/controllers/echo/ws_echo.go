package echo

import (
	"context"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

type WsEchoController struct {
	logger *logrus.Logger
}

func NewWsEchoController(
	logger *logrus.Logger,
) *WsEchoController {
	return &WsEchoController{
		logger,
	}
}

func (instance *WsEchoController) RelativePath() string {
	return "/echo"
}

func (instance *WsEchoController) Handle(ctx context.Context, conn *websocket.Conn) {
	defer conn.Close(websocket.StatusInternalError, "the sky is falling")

	for {
		// Read a message from the client
		_, message, err := conn.Read(ctx)
		if err != nil {
			break
		}

		instance.logger.WithContext(ctx).WithFields(logrus.Fields{
			"message": message,
		}).Debug("message received")

		// Echo the message back to the client
		err = conn.Write(ctx, websocket.MessageText, message)
		if err != nil {
			break
		}
	}
}
