package backdrop

import (
	"backend/types/payload"
	"backend/utils/value"
)

func GetBackdrop() *payload.BackdropState {
	return &payload.BackdropState{
		Time: value.Ptr("01:45"),
		Date: value.Ptr("Thu, May 25"),
		Whether: &payload.BackdropWeather{
			Icon:   value.Ptr("http://www.gstatic.com/images/icons/material/apps/weather/2x/clear_night_dark_color_96dp.png"),
			Status: value.Ptr("Clear Night"),
			Temp:   value.Ptr("24"),
		},
		CurrentEvent: &payload.BackdropEvent{
			Title: value.Ptr("Morning Break"),
			Time:  value.Ptr("10:00 AM - 10:30 AM"),
		},
		NextEvent: &payload.BackdropEvent{
			Title: value.Ptr("Hackathon Session"),
			Time:  value.Ptr("10:30 AM - 12:00 PM"),
		},
		NowPlaying: &payload.BackdropNowPlaying{
			CoverURL: value.Ptr("https://lh3.googleusercontent.com/3ABWzObIuJPlSmQ5K9WWbn0UV4i2gyPphxACYCA0PE50BYCsPn4w8b90JjNYbsUiPqSi-o49oRxyp-aj=w544-h544-l90-rj"),
			Title:    value.Ptr("Burn"),
			Artist:   value.Ptr("Ellie Goulding"),
			Album:    value.Ptr("Halcyon Days"),
			QueueBy:  value.Ptr("APISIT MANEERAT"),
		},
	}
}
