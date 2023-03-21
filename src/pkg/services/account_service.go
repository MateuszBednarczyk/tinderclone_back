package services

import (
	"tinderclone_back/src/pkg/stores"
	"tinderclone_back/src/pkg/utils"
)

type IAccounter interface {
	GetAccountInformations(username string) *Result
}

type accounter struct {
	userStore stores.IUserStore
	userUtil  utils.IUserUtil
}

func NewAccounter(store stores.IUserStore, userUtil utils.IUserUtil) *accounter {
	return &accounter{
		userStore: store,
		userUtil:  userUtil,
	}
}

func (s *accounter) GetAccountInformations(username string) *Result {
	user, err := s.userStore.SelectUserByUsername(username)
	if err != nil {
		return CreateServiceResult("Couldn't find an user", 404, []interface{}{})
	}
	response := s.userUtil.ProcessUserEntityToDTO(user)

	return CreateServiceResult("User found", 200, []interface{}{response})
}
