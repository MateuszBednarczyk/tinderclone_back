package services

import "tinderclone_back/src/pkg/stores"

type IAccounter interface {
	GetAccountInformations(username string) *Result
}

type accounter struct {
	userStore stores.IUserStore
}

func NewAccounter(store stores.IUserStore) *accounter {
	return &accounter{
		userStore: store,
	}
}

func (s *accounter) GetAccountInformations(username string) *Result {
	user, err := s.userStore.SelectUserByUsername(username)
	if err != nil {
		return CreateServiceResult("Couldn't find an user", 404, []interface{}{})
	}

	return CreateServiceResult("User found", 200, []interface{}{user})
}
