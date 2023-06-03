package payload

type BackdropState struct {
	Time         *string             `json:"time"`
	Date         *string             `json:"date"`
	Whether      *BackdropWeather    `json:"whether"`
	CurrentEvent *BackdropEvent      `json:"current_event"`
	NextEvent    *BackdropEvent      `json:"next_event"`
	NowPlaying   *BackdropNowPlaying `json:"now_playing"`
}

type BackdropEvent struct {
	Title *string `json:"title"`
	Time  *string `json:"time"`
}

type BackdropNowPlaying struct {
	QueueId  *uint64 `json:"queueId"`
	CoverURL *string `json:"cover_url"`
	Title    *string `json:"title"`
	Artist   *string `json:"artist"`
	Album    *string `json:"album"`
	QueueBy  *string `json:"queue_by"`
}

type BackdropWeather struct {
	Icon   *string  `json:"icon"`
	Status *string  `json:"status"`
	Temp   *float64 `json:"temp"`
	Aqi    *float64 `json:"aqi"`
	Uv     *float64 `json:"uv"`
}

// KMUTT LAT, LONG - 13.652162312808013, 100.4963937840923
