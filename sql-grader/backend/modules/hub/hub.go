package ihub

import (
	"database/sql"

	"gorm.io/gorm"

	"backend/types/extern"
)

var b *extern.Base
var h *extern.Hub

func Init(sqlDB *sql.DB, DB *gorm.DB) *extern.Hub {
	h = &extern.Hub{
		Sessions:         make(map[uint64]*extern.Session),
		SessionDbNameMap: make(map[string]*extern.Session),
		Mocks:            make(map[uint64]*extern.Mock),
	}

	b = &extern.Base{
		SqlDB:        sqlDB,
		DB:           DB,
		FirebaseApp:  nil,
		FirebaseAuth: nil,
		Fiber:        nil,
	}

	defer func() {
		// GradeSchedule()
	}()

	return h
}
