package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;auto_increment;uniqueIndex;not null" json:"ID"`
	Username string    `gorm:"uniqueIndex;not null" json:"Username"`
	Password string    `gorm:"not null" json:"Password"`
}
