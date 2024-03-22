package model

type UserConferenceRole struct {
	UserID       int    `gorm:"primaryKey"`
	ConferenceID int    `gorm:"primaryKey"`
	Role         string `gorm:"primaryKey"`
}
