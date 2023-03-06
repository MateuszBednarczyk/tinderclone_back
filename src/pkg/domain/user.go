package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"Id"`
	Username string    `gorm:"uniqueIndex;not null;type:varchar(25)" json:"Username"`
	Password string    `gorm:"not null" json:"Password"`
	Name     string    `gorm:"not null;type:varchar(25)" json:"Name"`
	Surname  string    `gorm:"not null;type:varchar(30)" json:"Surname"`
	Country  string    `gorm:"not null;type:varchar(25)" json:"Country"`
	City     string    `gorm:"not null;type:varchar(25)" json:"City"`
}
