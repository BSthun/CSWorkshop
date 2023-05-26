package payload

type EventData struct {
	Events []*EventItem `json:"events"`
}

type EventItem struct {
	Name      *string `yaml:"name"`
	StartTime *string `yaml:"start_time"`
	EndTime   *string `yaml:"end_time"`
}
