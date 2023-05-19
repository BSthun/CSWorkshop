package model

type UserTrackLike struct {
	UserId  *uint64 `gorm:"primaryKey"`
	User    *User   `gorm:"foreignKey:UserId"`
	TrackId *uint64 `gorm:"primaryKey"`
	Track   *Track  `gorm:"foreignKey:TrackId"`
}
