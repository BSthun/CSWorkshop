package websocket

import (
	"backend/modules"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/types/extern"
)

func ServeUpdate(conn *websocket.Conn) {
	// * Increment connection count
	modules.Hub.ConnectionIncrementMutex.Lock()
	modules.Hub.ConnectionIncrement++
	modules.Hub.Connections[modules.Hub.ConnectionIncrement] = conn
	modules.Hub.ConnectionIncrementMutex.Unlock()

	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			break
		}

		_ = conn.WriteJSON(&extern.OutboundMessage{
			Event:   extern.EchoEvent,
			Payload: p,
		})
	}

	// * Close connection
	if err := conn.Close(); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE: " + err.Error())
	}
}
