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
	modules.Hub.BackdropClientConnectionIncrementMutex.Lock()
	modules.Hub.BackdropClientConnectionIncrement++
	modules.Hub.BackdropClientConnections[modules.Hub.BackdropClientConnectionIncrement] = conn
	increment := modules.Hub.BackdropClientConnectionIncrement
	modules.Hub.BackdropClientConnectionIncrementMutex.Unlock()

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
	modules.Hub.BackdropClientConnectionIncrementMutex.Lock()
	delete(modules.Hub.BackdropClientConnections, increment)
	modules.Hub.BackdropClientConnectionIncrementMutex.Unlock()
}
