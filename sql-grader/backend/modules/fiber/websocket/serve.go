package websocket

import (
	"strconv"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/modules"
	"backend/types/extern"
	"backend/utils/value"
)

func Serve(conn *websocket.Conn) {
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
	session, ok := modules.Hub.Sessions[eid]
	if !ok {
		_ = conn.WriteJSON(&extern.OutboundMessage{
			Event:   extern.ErrorEvent,
			Payload: "Enrollment session not found",
		})
		return
	}
	if *session.Token != token {
		_ = conn.WriteJSON(&extern.OutboundMessage{
			Event:   extern.ErrorEvent,
			Payload: "Invalid token",
		})
		return
	}

	// * Check if connection already exists
	if session.Conn != nil {
		HandleConnectionSwitch(session, conn)
	}

	// * Assign connection
	session.ConnMutex.Lock()
	session.Conn = conn
	session.ConnMutex.Unlock()

	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			break
		}

		session.Emit(&extern.OutboundMessage{
			Event:   extern.EchoEvent,
			Payload: p,
		})
	}

	// * Close connection
	if err := conn.Close(); err != nil {
		logrus.Warn("UNHANDLED CONNECTION CLOSE: " + err.Error())
	}

	// * Reset player connection
	session.Conn = nil

	// * Unlock in case of connection switch
	if value.MutexLocked(session.ConnMutex) {
		session.ConnMutex.Unlock()
	} else {
		// * Session completely closed
		delete(modules.Hub.Sessions, eid)
	}
}
