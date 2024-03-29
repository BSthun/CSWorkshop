package websocket

import (
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	ihub "backend/modules/hub"
	"backend/types/extern"
)

func HandleConnectionSwitch(state *ihub.Session) {
	// * Connection switch
	logrus.Warn("CONNECTION SWITCH")
	state.Emit(&extern.OutboundMessage{
		Event:   extern.ConnectionSwitchEvent,
		Payload: nil,
	})

	state.ConnMutex.Lock()
	if err := state.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE MESSAGE: " + err.Error())
	}
	if err := state.Conn.Close(); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE: " + err.Error())
	}
}

func HandleMockConnectionSwitch(mock *ihub.Mock) {
	mock.ConnMutex.Lock()
	if err := mock.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE MESSAGE: " + err.Error())
	}
	if err := mock.Conn.Close(); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE: " + err.Error())
	}
}
