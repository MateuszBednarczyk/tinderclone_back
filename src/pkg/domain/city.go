package domain

import (
	"github.com/google/uuid"
)

type City struct {
	CityID    uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"Id"`
	CityName  string    `gorm:"uniqueIndex;type:varchar(40);not null" json:"Name"`
	CountryID uuid.UUID
	Country   Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Users     []User  `gorm:"foreignKey:UserID"`
}
