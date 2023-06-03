package hubModel

type Hub struct {
	MusicClient       *Client            `json:"-"`
	BackdropClient    *Client            `json:"-"`
	SpotifyCredential *SpotifyCredential `json:"-"`
}
