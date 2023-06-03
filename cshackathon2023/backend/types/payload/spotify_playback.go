package payload

type SpotifyCredentials struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	TokenType    *string `json:"token_type"`
	ExpiresIn    *int64  `json:"expires_in"`
	Scope        *string `json:"scope"`
	Error        *string `json:"error"`
}

type SpotifyPlaybackState struct {
	Device               *SpotifyDevice  `json:"device"`
	ShuffleState         *bool           `json:"shuffle_state"`
	RepeatState          *string         `json:"repeat_state"`
	Timestamp            *int64          `json:"timestamp"`
	Context              *SpotifyContext `json:"context"`
	ProgressMs           *int64          `json:"progress_ms"`
	Item                 *SpotifyTrack   `json:"item"`
	CurrentlyPlayingType *string         `json:"currently_playing_type"`
	IsPlaying            *bool           `json:"is_playing"`
}

type SpotifyDevice struct {
	Id               *string `json:"id"`
	IsActive         *bool   `json:"is_active"`
	IsPrivateSession *bool   `json:"is_private_session"`
	IsRestricted     *bool   `json:"is_restricted"`
	Name             *string `json:"name"`
	Type             *string `json:"type"`
	VolumePercent    *int    `json:"volume_percent"`
}

type SpotifyExternalIds struct {
	Isrc string `json:"isrc"`
}

type SpotifyContext struct {
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Href         string               `json:"href"`
	Type         string               `json:"type"`
	Uri          string               `json:"uri"`
}

type SpotifyTrack struct {
	Album        *SpotifyAlbum        `json:"album"`
	Artists      []*SpotifyArtist     `json:"artists"`
	DiscNumber   *int64               `json:"disc_number"`
	DurationMs   *int64               `json:"duration_ms"`
	Explicit     *bool                `json:"explicit"`
	ExternalIds  *SpotifyExternalIds  `json:"external_ids"`
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Href         *string              `json:"href"`
	Id           *string              `json:"id"`
	IsLocal      *bool                `json:"is_local"`
	Name         *string              `json:"name"`
	Popularity   *int64               `json:"popularity"`
	PreviewUrl   *string              `json:"preview_url"`
	TrackNumber  *int64               `json:"track_number"`
	Type         *string              `json:"type"`
	Uri          *string              `json:"uri"`
}

type SpotifyAlbum struct {
	AlbumType            *string              `json:"album_type"`
	Artists              []*SpotifyArtist     `json:"artists"`
	AvailableMarkets     []string             `json:"available_markets"`
	ExternalUrls         *SpotifyExternalUrls `json:"external_urls"`
	Href                 *string              `json:"href"`
	Id                   *string              `json:"id"`
	Images               []*SpotifyImage      `json:"images"`
	Name                 *string              `json:"name"`
	ReleaseDate          *string              `json:"release_date"`
	ReleaseDatePrecision *string              `json:"release_date_precision"`
	TotalTracks          *int                 `json:"total_tracks"`
	Type                 *string              `json:"type"`
	Uri                  *string              `json:"uri"`
}

type SpotifyArtist struct {
	ExternalUrls *SpotifyExternalUrls `json:"external_urls"`
	Href         *string              `json:"href"`
	Id           *string              `json:"id"`
	Name         *string              `json:"name"`
	Type         *string              `json:"type"`
	Uri          *string              `json:"uri"`
}

type SpotifyTrackItem struct {
	Items []*SpotifyTrack
}

type SpotifyTrackSearch struct {
	Tracks *SpotifyTrackItem
}
