package payload

import "time"

type EventData struct {
	Events []*EventItem `yaml:"events"`
}

type EventItem struct {
	Name      *string    `yaml:"name"`
	StartTime *time.Time `yaml:"start_time"`
	EndTime   *time.Time `yaml:"end_time"`
}
