package backdrop

import (
	"backend/types/payload"
	"backend/utils/value"
)

func GetNowPlaying() *payload.BackdropNowPlaying {
	return &payload.BackdropNowPlaying{
		CoverURL: value.Ptr("https://lh3.googleusercontent.com/3ABWzObIuJPlSmQ5K9WWbn0UV4i2gyPphxACYCA0PE50BYCsPn4w8b90JjNYbsUiPqSi-o49oRxyp-aj=w544-h544-l90-rj"),
		Title:    value.Ptr("Burn"),
		Artist:   value.Ptr("Ellie Goulding"),
		Album:    value.Ptr("Halcyon Days"),
		QueueBy:  value.Ptr("APISIT MANEERAT"),
	}
}
