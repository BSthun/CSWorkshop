package websocket

import (
	"strconv"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/modules"
	"backend/types/extern"
	"backend/utils/value"
)

func ServeMock(conn *websocket.Conn) {
	// * Parse query parameters
	eid, err := strconv.ParseUint(conn.Query("eid"), 10, 64)
	if err != nil {
		_ = conn.WriteJSON(&extern.OutboundMessage{
			Event:   extern.ErrorEvent,
			Payload: "Unable to parse enrollment ID",
		})
		return
	}
	token := conn.Query("token")

	// * Fetch enrollment session
	mock, ok := modules.Hub.Mocks[eid]
	if !ok {
		_ = conn.WriteJSON(&extern.OutboundMessage{
			Event:   extern.ErrorEvent,
			Payload: "Enrollment session not found",
		})
		return
	}
	if *mock.Token != token {
		_ = conn.WriteJSON(&extern.OutboundMessage{
			Event:   extern.ErrorEvent,
			Payload: "Invalid token",
		})
		return
	}

	if mock.Conn != nil {
		HandleMockConnectionSwitch(mock)
	}

	// * Assign connection
	mock.ConnMutex.Lock()
	mock.Conn = conn
	mock.ConnMutex.Unlock()

	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			break
		}

		_ = mock.Conn.WriteMessage(t, p)
	}

	// * Close connection
	if err := conn.Close(); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE: " + err.Error())
	}

	// * Reset player connection
	mock.Conn = nil

	// * Unlock in case of connection switch
	if value.MutexLocked(mock.ConnMutex) {
		mock.ConnMutex.Unlock()
	}
}
