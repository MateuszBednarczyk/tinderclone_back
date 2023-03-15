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
	UserID   uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"Id"`
	Username string    `gorm:"uniqueIndex;not null;type:varchar(25)" json:"Username"`
	Password string    `gorm:"not null" json:"Password"`
	Name     string    `gorm:"not null;type:varchar(25)" json:"Name"`
	Surname  string    `gorm:"not null;type:varchar(30)" json:"Surname"`
	Country  uuid.UUID `gorm:"type:uuid;foreignKey:CountryIDconstraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Country"`
	City     uuid.UUID `gorm:"type:uuid;foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"City"`
	Role     Role      `gorm:"type:int" json:"Role"`
}
