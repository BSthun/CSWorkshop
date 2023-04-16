package ihub

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/modules/config"
	"backend/types/embed"
	"backend/types/extern"
)

type Hub struct {
	Sessions map[uint64]*Session
}

type Session struct {
	Id         *uint64           `json:"id"`
	Credential *embed.Credential `json:"credential"`
	DbName     *string           `json:"db_name"`
	Token      *string           `json:"token"`
	Conn       *websocket.Conn   `json:"conn"`
	ConnMutex  *sync.Mutex       `json:"conn_mutex"`
}

func (r *Session) Emit(payload *extern.OutboundMessage) {
	if r.Conn == nil || r.Conn.Conn == nil {
		return
	}

	r.ConnMutex.Lock()
	if err := r.Conn.WriteJSON(payload); err != nil {
		logrus.Warn(fmt.Sprintf("WRITING MESSAGE FAILURE: %s", err.Error()))
	}
	r.ConnMutex.Unlock()
}

func Init(conf *iconfig.Config) *Hub {
	hub := &Hub{
		Sessions: make(map[uint64]*Session),
	}

	return hub
}
