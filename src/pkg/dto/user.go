package dto

import (
	"github.com/google/uuid"

	"tinderclone_back/src/pkg/domain"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Name      string
	Surname   string
	Countries []string
	Cities    []string
	Role      domain.Role
}
