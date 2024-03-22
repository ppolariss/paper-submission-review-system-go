package model

type User struct {
	ID              int
	Username        string `gorm:"uniqueIndex"`
	RealName        string
	Password        string
	Email           string
	InstitutionName string
	Area            string
}
