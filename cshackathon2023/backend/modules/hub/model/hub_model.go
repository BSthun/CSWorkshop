package hubModel

import (
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Hub struct {
	MusicClientConnections                 map[int64]*websocket.Conn `json:"connections"`
	MusicClientConnectionIncrement         int64                     `json:"musicClientConnectionIncrement"`
	MusicClientConnectionIncrementMutex    *sync.Mutex               `json:"-"`
	BackdropClientConnections              map[int64]*websocket.Conn `json:"backdropClientConnections"`
	BackdropClientConnectionIncrement      int64                     `json:"backdropClientConnectionIncrement"`
	BackdropClientConnectionIncrementMutex *sync.Mutex               `json:"-"`
}
