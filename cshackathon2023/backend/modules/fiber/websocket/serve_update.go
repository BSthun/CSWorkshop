package websocket

import (
	"backend/modules"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/types/extern"
)

func ServeMusicState(conn *websocket.Conn) {
	// * Increment connection count
	modules.Hub.MusicClient.IncrementMutex.Lock()
	modules.Hub.MusicClient.Increment++
	modules.Hub.MusicClient.Connections[modules.Hub.MusicClient.Increment] = conn
	increment := modules.Hub.MusicClient.Increment
	modules.Hub.MusicClient.IncrementMutex.Unlock()

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
	modules.Hub.MusicClient.WriteMutex.Lock()
	delete(modules.Hub.MusicClient.Connections, increment)
	modules.Hub.MusicClient.WriteMutex.Unlock()
}
