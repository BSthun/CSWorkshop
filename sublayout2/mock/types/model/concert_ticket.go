package model

type ConcertTicket struct {
	Id        *uint64  `gorm:"primaryKey;autoIncrement"`
	ConcertId *uint64  `gorm:"not null"`
	Concert   *Concert `gorm:"foreignKey:ConcertId"`
	Name      *string  `gorm:"type:varchar(256)"`
	CoverUrl  *string  `gorm:"not null"`
	Link      *string  `gorm:"not null"`
}
