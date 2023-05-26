package ihub

import (
	hubModel "backend/modules/hub/model"
)

var Hub *hubModel.Hub

func Init() *hubModel.Hub {
	// * Initialize hub
	Hub = &hubModel.Hub{
		MusicClient:       hubModel.NewClient(),
		BackdropClient:    hubModel.NewClient(),
		SpotifyCredential: hubModel.NewSpotifyCredential(),
	}

	// * Initialize cron jobs
	Cron()

	return Hub
}
