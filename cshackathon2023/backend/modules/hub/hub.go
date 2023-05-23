package ihub

import (
	hubModel "backend/modules/hub/model"
	"github.com/gofiber/websocket/v2"
	"sync"
)

func Init() *hubModel.Hub {
	return &hubModel.Hub{
		Connections:              make(map[int64]*websocket.Conn),
		ConnectionIncrement:      0,
		ConnectionIncrementMutex: new(sync.Mutex),
	}
}
