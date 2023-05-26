package backdrop

import (
	"backend/types/payload"
)

func GetBackdrop() *payload.BackdropState {
	date, time := GetDateTime()
	weather := GetWeather()
	currentEvent, nextEvent := GetEvent()
	nowPlaying := GetNowPlaying()

	return &payload.BackdropState{
		Time:         time,
		Date:         date,
		Whether:      weather,
		CurrentEvent: currentEvent,
		NextEvent:    nextEvent,
		NowPlaying:   nowPlaying,
	}
}
