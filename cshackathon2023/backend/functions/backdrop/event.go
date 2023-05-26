package backdrop

import (
	"backend/types/payload"
	"backend/utils/value"
)

func GetEvent() (*payload.BackdropEvent, *payload.BackdropEvent) {
	currentEvent := &payload.BackdropEvent{
		Title: value.Ptr("Morning Break"),
		Time:  value.Ptr("10:00 AM - 10:30 AM"),
	}

	nextEvent := &payload.BackdropEvent{
		Title: value.Ptr("Hackathon Session"),
		Time:  value.Ptr("10:30 AM - 12:00 PM"),
	}

	return currentEvent, nextEvent
}
