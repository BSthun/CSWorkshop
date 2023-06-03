package backdrop

import (
	"backend/types/payload"
	"backend/utils/value"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func GetEvent() (*payload.BackdropEvent, *payload.BackdropEvent) {
	events := new(payload.EventData)

	yml, err := os.ReadFile("resources/data/events.yaml")
	if err != nil {
		logrus.Fatal("UNABLE TO READ")
	}
	if err := yaml.Unmarshal(yml, events); err != nil {
		logrus.Fatal("UNABLE TO PARSE")
	}
	var currentEvent *payload.EventItem
	var nextEvent *payload.EventItem

	for i, s := range events.Events {
		if s.StartTime.Before(time.Now()) && s.EndTime.After(time.Now()) {
			currentEvent = s
			nextEvent = events.Events[i+1]
		}
	}

	//spew.Dump(events)
	stc := currentEvent.StartTime.Local().Format("15:04 PM")
	etc := currentEvent.EndTime.Local().Format("15:04 PM")
	showEvent := &payload.BackdropEvent{
		Title: currentEvent.Name,
		Time:  value.Ptr(stc + " - " + etc),
	}
	stn := nextEvent.StartTime.Local().Format("15:04 PM")
	etn := nextEvent.EndTime.Local().Format("15:04 PM")
	laterEvent := &payload.BackdropEvent{
		Title: nextEvent.Name,
		Time:  value.Ptr(stn + " - " + etn),
	}

	return showEvent, laterEvent
}
