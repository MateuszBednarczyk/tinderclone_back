package services

import (
	"golang.org/x/crypto/bcrypt"

	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
	"tinderclone_back/src/pkg/utils"
)

type IAuthenticator interface {
	LoginUser(dto dto.Credentials) *Result
}

type authenticator struct {
	userStore stores.IUserStore
	userUtil  utils.IUserUtil
}

func NewAuthenticator(store stores.IUserStore, userUtil utils.IUserUtil) *authenticator {
	return &authenticator{
		userStore: store,
		userUtil:  userUtil,
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
	tokens := Tokenizer().GenerateTokens(*s.userUtil.ProcessUserEntityToDTO(foundUser))

	return CreateServiceResult("Logged in", 200, tokens.Content)
}
