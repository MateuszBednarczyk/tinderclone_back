package domain

import (
	"github.com/google/uuid"
)

type Country struct {
	CountryID   uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"Id"`
	CountryName string    `gorm:"uniqueIndex;type:varchar(40);not null" json:"Name"`
	Cities      []City    `gorm:"type:uuid;foreignKey:CityID" json:"Cities"`
}
