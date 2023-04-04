package branchModel

type Lv3AlbumArtist struct {
	AlbumId  *uint64    `gorm:"primaryKey"`
	Album    *Lv3Album  `gorm:"foreignKey:AlbumId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ArtistId *uint64    `gorm:"primaryKey"`
	Artist   *Lv3Artist `gorm:"foreignKey:ArtistId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
