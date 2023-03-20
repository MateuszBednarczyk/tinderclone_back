package dto

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Username  string
	Name      string
	Surname   string
	Countries []string
	Cities    []string
}
