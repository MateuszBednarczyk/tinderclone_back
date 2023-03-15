package domain

import (
	"github.com/google/uuid"
)

type Role int

const (
	admin Role = iota
	moderator
	user
)

type User struct {
	UserID    uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"Id"`
	Username  string    `gorm:"uniqueIndex;not null;type:varchar(25)" json:"Username"`
	Password  string    `gorm:"not null" json:"Password"`
	Name      string    `gorm:"not null;type:varchar(25)" json:"Name"`
	Surname   string    `gorm:"not null;type:varchar(30)" json:"Surname"`
	CountryID uuid.UUID
	Country   Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CityID    uuid.UUID
	City      City `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role      Role `gorm:"type:int" json:"Role"`
}
