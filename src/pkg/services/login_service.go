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
	userStore stores.IUserStore
}

func NewAuthenticator(store stores.IUserStore) *authenticator {
	return &authenticator{
		userStore: store,
	}
}

func (s *authenticator) LoginUser(requestBody dto.Credentials) *Result {
	var err error

	foundUser, err := s.userStore.SelectUserByUsername(requestBody.Username)
	if err != nil {
		return CreateServiceResult("Couldn't find user", 404, []interface{}{err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(requestBody.Password))
	if err != nil {
		return CreateServiceResult("Bad credentials", 401, []interface{}{})
	}

	tokens := Tokenizer().GenerateTokens(*foundUser)

	return CreateServiceResult("Logged in", 200, tokens.Content)
}
