package extern

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/types/payload"
)

type Hub struct {
	Sessions         map[uint64]*Session
	SessionDbNameMap map[string]*Session
	Mocks            map[uint64]*Mock
}

type Session struct {
	Id          *uint64                            `json:"id"`
	LabId       *uint64                            `json:"labId"`
	UserId      *uint64                            `json:"userId"`
	Db          *sql.DB                            `json:"db"`
	DbName      *string                            `json:"dbName"`
	DbValid     *bool                              `json:"dbValid"`
	Token       *string                            `json:"token"`
	CurrentTask *uint64                            `json:"currentTask"`
	TaskResults map[uint64]*payload.LabStateResult `json:"taskResults"`
	Conn        *websocket.Conn                    `json:"conn"`
	ConnMutex   *sync.Mutex                        `json:"connMutex"`
}

func (r *Session) Emit(payload *OutboundMessage) {
	if r.Conn == nil || r.Conn.Conn == nil {
		return
	}

	r.ConnMutex.Lock()
	if err := r.Conn.WriteJSON(payload); err != nil {
		logrus.Warn(fmt.Sprintf("WRITING MESSAGE FAILURE: %s", err.Error()))
	}
	r.ConnMutex.Unlock()
}

type Mock struct {
	Lines     []string        `json:"lines"`
	Token     *string         `json:"token"`
	Conn      *websocket.Conn `json:"conn"`
	ConnMutex *sync.Mutex     `json:"conn_mutex"`
}

func (r *Mock) Append(line string) {
	r.Lines = append(r.Lines, line)
	if r.Conn != nil {
		r.ConnMutex.Lock()
		_ = r.Conn.WriteMessage(websocket.TextMessage, []byte(line))
		r.ConnMutex.Unlock()
	}
}
