package websocket

import (
	"backend/functions/backdrop"
	"backend/modules"
	"backend/types/extern"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

func ServeBackdropState(conn *websocket.Conn) {
	// * Increment connection count
	modules.Hub.BackdropClient.IncrementMutex.Lock()
	modules.Hub.BackdropClient.Increment++
	modules.Hub.BackdropClient.Connections[modules.Hub.BackdropClient.Increment] = conn
	increment := modules.Hub.BackdropClient.Increment
	modules.Hub.BackdropClient.IncrementMutex.Unlock()

	// * Send initial state
	_ = conn.WriteJSON(&extern.OutboundMessage{
		Event:   extern.BackdropUpdateEvent,
		Payload: backdrop.GetBackdrop(),
	})

	// * Listen for messages
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
	modules.Hub.BackdropClient.WriteMutex.Lock()
	delete(modules.Hub.MusicClient.Connections, increment)
	modules.Hub.BackdropClient.WriteMutex.Unlock()
}
