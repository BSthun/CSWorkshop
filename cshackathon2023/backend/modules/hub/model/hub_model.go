package hubModel

import (
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Hub struct {
	Connections              map[int64]*websocket.Conn `json:"connections"`
	ConnectionIncrement      int64                     `json:"connection_increment"`
	ConnectionIncrementMutex *sync.Mutex               `json:"-"`
}
