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
		return &Result{
			Message: "Username cannot be blank",
			Code:    400,
			Content: []interface{}{},
		}
	}

	if !isPasswordValid(dto.Password) {
		return &Result{
			Message: "Password cannot be blank",
			Code:    400,
			Content: []interface{}{},
		}
	}

	if stores.IsUsernameAlreadyTaken(dto.Username) {
		return &Result{
			Message: "Username is already taken",
			Code:    409,
			Content: []interface{}{},
		}
	}

	hash, err := hashPassword(dto.Password)
	if err != nil {
		return &Result{
			Message: "Couldn't hash password",
			Code:    500,
			Content: []interface{}{},
		}
	}

	user := &domain.User{
		Username: dto.Username,
		Password: string(hash),
	}

	result := stores.SaveUser(user)
	if result != nil {
		return &Result{
			Message: "There was an error, while attempt of saving user",
			Code:    500,
			Content: []interface{}{},
		}
	}

	return &Result{
		Message: "Account created",
		Code:    200,
		Content: []interface{}{},
	}
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
