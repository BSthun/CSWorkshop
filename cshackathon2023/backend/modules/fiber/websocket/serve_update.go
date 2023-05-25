package websocket

import (
	"backend/modules"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/types/extern"
)

func ServeMusicState(conn *websocket.Conn) {
	// * Increment connection count
	modules.Hub.MusicClientConnectionIncrementMutex.Lock()
	modules.Hub.MusicClientConnectionIncrement++
	modules.Hub.MusicClientConnections[modules.Hub.MusicClientConnectionIncrement] = conn
	increment := modules.Hub.MusicClientConnectionIncrement
	modules.Hub.MusicClientConnectionIncrementMutex.Unlock()

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

	// * Delete connection
	modules.Hub.MusicClientConnectionIncrementMutex.Lock()
	delete(modules.Hub.MusicClientConnections, increment)
	modules.Hub.MusicClientConnectionIncrementMutex.Unlock()
}
