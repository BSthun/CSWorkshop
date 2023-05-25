package payload

type BackdropState struct {
	Time         *string             `json:"time,omitempty"`
	Date         *string             `json:"date,omitempty"`
	Whether      *BackdropWeather    `json:"whether,omitempty"`
	CurrentEvent *BackdropEvent      `json:"current_event,omitempty"`
	NextEvent    *BackdropEvent      `json:"next_event,omitempty"`
	NowPlaying   *BackdropNowPlaying `json:"now_playing,omitempty"`
}

type BackdropEvent struct {
	Title *string `json:"title,omitempty"`
	Time  *string `json:"time,omitempty"`
}

type BackdropNowPlaying struct {
	CoverURL *string `json:"cover_url,omitempty"`
	Title    *string `json:"title,omitempty"`
	Artist   *string `json:"artist,omitempty"`
	Album    *string `json:"album,omitempty"`
	QueueBy  *string `json:"queue_by,omitempty"`
}

type BackdropWeather struct {
	Icon   *string `json:"icon,omitempty"`
	Status *string `json:"status,omitempty"`
	Temp   *string `json:"temp,omitempty"`
}
