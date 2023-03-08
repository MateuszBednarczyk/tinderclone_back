package domain

import (
	"github.com/google/uuid"
)

type City struct {
	CityID   uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"Id"`
	CityName string    `gorm:"type:varchar(40);not null" json:"Name"`
}
