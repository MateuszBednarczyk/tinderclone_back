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

func NewAuthorizer() *authorizer {
	return &authorizer{}
}

func (s *authorizer) LoginUser(requestBody dto.Credentials) *Result {
	var err error

	foundUser, err := stores.SelectUserByUsername(requestBody.Username)
	if err != nil {
		return CreateServiceResult("Couldn't find user", 404, []interface{}{err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(requestBody.Password))
	if err != nil {
		return CreateServiceResult("Bad credentials", 401, []interface{}{})
	}
	tokens := JwtService().GenerateTokens(requestBody.Username)

	return CreateServiceResult("Logged in", 200, tokens.Content)
}
