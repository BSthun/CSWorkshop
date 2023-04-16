package websocket

import (
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	ihub "backend/modules/hub"
	"backend/types/extern"
)

func HandleConnectionSwitch(state *ihub.Session, conn *websocket.Conn) {
	// * Connection switch
	logrus.Warn("CONNECTION SWITCH")
	state.Emit(&extern.OutboundMessage{
		Event:   extern.ConnectionSwitchEvent,
		Payload: nil,
	})

	state.ConnMutex.Lock()
	if err := conn.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE MESSAGE: " + err.Error())
	}
	if err := conn.Conn.Close(); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE: " + err.Error())
	}
}
