package domain

import (
	"github.com/google/uuid"
)

type Country struct {
	CountryID   uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"d"`
	CountryName string    `gorm:"uniqueIndex;type:varchar(40);not null" json:"name"`
}
