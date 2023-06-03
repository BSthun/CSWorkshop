package backdrop

import (
	"backend/modules"
	"backend/types/payload"
	"backend/utils/value"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

var weatherData *payload.BackdropWeather
var weatherLastUpdated *time.Time

func GetWeather() *payload.BackdropWeather {
	if weatherData == nil || time.Since(*weatherLastUpdated) > 2*time.Minute {
		weatherData = FetchWeather()
		weatherLastUpdated = value.Ptr(time.Now())
	}
	return weatherData
}

func FetchWeather() *payload.BackdropWeather {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	// * Send request
	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json?key="+modules.Conf.WeatherApiKey+"&aqi=yes&q=13.651784878173968,100.4964698106748", nil)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)

	var createQrResponse *payload.Weather
	err = json.NewDecoder(res.Body).Decode(&createQrResponse)
	if err != nil {
		return &payload.BackdropWeather{
			Icon:   value.Ptr("https://static.vecteezy.com/system/resources/previews/017/172/383/original/warning-message-concept-represented-by-exclamation-mark-icon-exclamation-symbol-in-circle-png.png"),
			Status: value.Ptr("Error"),
			Temp:   value.Ptr(0.0),
			Aqi:    value.Ptr(0.0),
		}
	}

	return &payload.BackdropWeather{
		Icon:   value.Ptr("https:" + *createQrResponse.Current.Condition.Icon),
		Status: createQrResponse.Current.Condition.Text,
		Temp:   createQrResponse.Current.TempC,
		Aqi:    createQrResponse.Current.AirQuality.Pm2_5,
	}
}
