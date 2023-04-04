package model

type Client struct {
	Id                  *uint64 `gorm:"primaryKey"`
	SpotifyClientId     *string `gorm:"type:varchar(255); index:spotify_client_id,unique; not null" yaml:"cid"`
	SpotifyClientSecret *string `gorm:"type:varchar(255); not null" yaml:"secret"`
	Authorization       *string `gorm:"type:varchar(255); not null"`
}
