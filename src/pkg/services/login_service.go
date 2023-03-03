package services

import (
	"golang.org/x/crypto/bcrypt"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

type ILoginService interface {
	LoginUser(dto dto.Credentials) *Result
}

type loginService struct {
}

func NewLoginService() *loginService {
	return &loginService{}
}

func (s *loginService) LoginUser(dto dto.Credentials) *Result {
	var err error

	foundUser, err := stores.SelectUserByUsername(dto.Username)
	if err != nil {
		return NewResult("Couldn't find user", 404, []interface{}{err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(dto.Password))
	if err != nil {
		return NewResult("Bad credentials", 401, []interface{}{})
	}

	return NewResult("Logged in", 200, []interface{}{"tokens"})
}
