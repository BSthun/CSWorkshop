package ihub

import (
	hubModel "backend/modules/hub/model"
)

var Hub *hubModel.Hub

func Init() *hubModel.Hub {
	// * Initialize hub
	Hub = &hubModel.Hub{
		MusicClient:    hubModel.NewClient(),
		BackdropClient: hubModel.NewClient(),
	}

	// * Initialize cron jobs
	Cron()

	return Hub
}
