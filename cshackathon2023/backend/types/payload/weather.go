package payload

type Weather struct {
	Location *Location `json:"location,omitempty"`
	Current  *Current  `json:"current,omitempty"`
}

type Current struct {
	LastUpdatedEpoch *int64            `json:"last_updated_epoch,omitempty"`
	LastUpdated      *string           `json:"last_updated,omitempty"`
	TempC            *float64          `json:"temp_c,omitempty"`
	TempF            *float64          `json:"temp_f,omitempty"`
	IsDay            *int64            `json:"is_day,omitempty"`
	Condition        *Condition        `json:"condition,omitempty"`
	WindMph          *float64          `json:"wind_mph,omitempty"`
	WindKph          *float64          `json:"wind_kph,omitempty"`
	WindDegree       *int64            `json:"wind_degree,omitempty"`
	WindDir          *string           `json:"wind_dir,omitempty"`
	PressureMB       *float64          `json:"pressure_mb,omitempty"`
	PressureIn       *float64          `json:"pressure_in,omitempty"`
	PrecipMm         *float64          `json:"precip_mm,omitempty"`
	PrecipIn         *float64          `json:"precip_in,omitempty"`
	Humidity         *int64            `json:"humidity,omitempty"`
	Cloud            *int64            `json:"cloud,omitempty"`
	FeelslikeC       *float64          `json:"feelslike_c,omitempty"`
	FeelslikeF       *float64          `json:"feelslike_f,omitempty"`
	VisKM            *float64          `json:"vis_km,omitempty"`
	VisMiles         *float64          `json:"vis_miles,omitempty"`
	Uv               *float64          `json:"uv,omitempty"`
	GustMph          *float64          `json:"gust_mph,omitempty"`
	GustKph          *float64          `json:"gust_kph,omitempty"`
	AirQuality       *AirQualityDetail `json:"air_quality,omitempty"`
}

type AirQualityDetail struct {
	Co           *float64 `json:"co,omitempty"`
	No2          *float64 `json:"no2,omitempty"`
	O3           *float64 `json:"o3,omitempty"`
	So2          *float64 `json:"so2,omitempty"`
	Pm2_5        *float64 `json:"pm2_5,omitempty"`
	PM10         *float64 `json:"pm10,omitempty"`
	UsEpaIndex   *int64   `json:"us-epa-index,omitempty"`
	GbDefraIndex *int64   `json:"gb-defra-index,omitempty"`
}

type Condition struct {
	Text *string `json:"text,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Code *int64  `json:"code,omitempty"`
}

type Location struct {
	Name           *string  `json:"name,omitempty"`
	Region         *string  `json:"region,omitempty"`
	Country        *string  `json:"country,omitempty"`
	Lat            *float64 `json:"lat,omitempty"`
	Lon            *float64 `json:"lon,omitempty"`
	TzID           *string  `json:"tz_id,omitempty"`
	LocaltimeEpoch *int64   `json:"localtime_epoch,omitempty"`
	Localtime      *string  `json:"localtime,omitempty"`
}
