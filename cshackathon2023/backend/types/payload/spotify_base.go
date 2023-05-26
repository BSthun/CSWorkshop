package payload

type SpotifyExternalUrls struct {
	Spotify string `json:"spotify"`
}

type SpotifyFollowers struct {
	Href  *string `json:"href"`
	Total *int    `json:"total"`
}

type SpotifyImage struct {
	Height *int    `json:"height"`
	Url    *string `json:"url"`
	Width  *int    `json:"width"`
}

type SpotifyError struct {
	Error *string `json:"error"`
}

type SpotifyApiError struct {
	Error struct {
		Status  *int    `json:"status"`
		Message *string `json:"message"`
	} `json:"error"`
}
