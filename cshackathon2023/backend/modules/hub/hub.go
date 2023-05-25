package ihub

import (
	hubModel "backend/modules/hub/model"
	"github.com/gofiber/websocket/v2"
	"sync"
)

var Hub *hubModel.Hub

func Init() *hubModel.Hub {
	// * Initialize hub
	Hub = &hubModel.Hub{
		MusicClientConnections:                 make(map[int64]*websocket.Conn),
		MusicClientConnectionIncrement:         0,
		MusicClientConnectionIncrementMutex:    new(sync.Mutex),
		BackdropClientConnections:              make(map[int64]*websocket.Conn),
		BackdropClientConnectionIncrement:      0,
		BackdropClientConnectionIncrementMutex: new(sync.Mutex),
	}

	// * Initialize cron jobs
	Cron()

	return Hub
}
