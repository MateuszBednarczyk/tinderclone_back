package services

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

type IRegisterService interface {
	RegisterUser(dto dto.RegisterUser) *Result
}

type registerService struct {
}

func NewRegisterService() *registerService {
	return &registerService{}
}

func (s *registerService) RegisterUser(dto dto.RegisterUser) *Result {
	if !isUsernameValid(dto.Username) {
		return NewResult("Username cannot be blank", 400, []interface{}{})
	}

	if !isPasswordValid(dto.Password) {
		return NewResult("Password cannot be blank", 400, []interface{}{})
	}

	if stores.IsUsernameAlreadyTaken(dto.Username) {
		return NewResult("Username is already taken", 409, []interface{}{})
	}

	hash, err := hashPassword(dto.Password)
	if err != nil {
		return NewResult("Couldn't hash password", 500, []interface{}{})
	}

	user := &domain.User{
		Username: dto.Username,
		Password: string(hash),
	}

	result := stores.SaveUser(user)
	if result != nil {
		return NewResult("There was an error, while attempt of saving user", 500, []interface{}{})
	}

	return NewResult("Account created", 200, []interface{}{})
}

func isUsernameValid(username string) bool {
	return len(strings.TrimSpace(username)) > 0
}

func isPasswordValid(password string) bool {
	return len(strings.TrimSpace(password)) > 0
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
