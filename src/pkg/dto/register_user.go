package dto

type RegisterUser struct {
	Username  string   `json:"Username" validate:"required,gte=1,lte=25"`
	Password  string   `json:"Password" validate:"required"`
	Name      string   `json:"Name" validate:"required,gte=1,lte=25"`
	Surname   string   `json:"Surname" validate:"required,gte=1,lte=30"`
	Countries []string `json:"Countries" validate:"required,gte=1,lte=25"`
	Cities    []string `json:"Cities" validate:"required"`
}
