package services

import (
	"golang.org/x/crypto/bcrypt"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

type IAuthorizer interface {
	LoginUser(dto dto.Credentials) *Result
}

type authorizer struct {
}

func NewLoginService() *authorizer {
	return &authorizer{}
}

func (s *authorizer) LoginUser(requestBody dto.Credentials) *Result {
	var err error

	foundUser, err := stores.SelectUserByUsername(requestBody.Username)
	if err != nil {
		return NewResult("Couldn't find user", 404, []interface{}{err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(requestBody.Password))
	if err != nil {
		return NewResult("Bad credentials", 401, []interface{}{})
	}
	tokens := JwtService().GenerateTokens(requestBody.Username)

	return NewResult("Logged in", 200, tokens.Content)
}
