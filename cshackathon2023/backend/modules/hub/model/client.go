package hubModel

import (
	"backend/types/extern"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Client struct {
	Connections    map[int64]*websocket.Conn `json:"connections"`
	Increment      int64                     `json:"increment"`
	IncrementMutex *sync.Mutex               `json:"-"`
	WriteMutex     *sync.Mutex               `json:"-"`
}

func NewClient() *Client {
	return &Client{
		Connections:    make(map[int64]*websocket.Conn),
		Increment:      0,
		IncrementMutex: new(sync.Mutex),
		WriteMutex:     new(sync.Mutex),
	}
}

func (r *Client) Emit(message *extern.OutboundMessage) {
	r.WriteMutex.Lock()
	defer r.WriteMutex.Unlock()

	for _, conn := range r.Connections {
		_ = conn.WriteJSON(message)
	}
}
