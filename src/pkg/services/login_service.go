package services

import (
	"golang.org/x/crypto/bcrypt"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

type IAuthenticator interface {
	LoginUser(dto dto.Credentials) *Result
}

type authenticator struct {
}

func NewAuthenticator() *authenticator {
	return &authenticator{}
}

func (s *authenticator) LoginUser(requestBody dto.Credentials) *Result {
	var err error

	foundUser, err := stores.SelectUserByUsername(requestBody.Username)
	if err != nil {
		return CreateServiceResult("Couldn't find user", 404, []interface{}{err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(requestBody.Password))
	if err != nil {
		return CreateServiceResult("Bad credentials", 401, []interface{}{})
	}

	isAdmin := foundUser.Role == 0
	tokens := Tokenizer().GenerateTokens(requestBody.Username, isAdmin)

	return CreateServiceResult("Logged in", 200, tokens.Content)
}
