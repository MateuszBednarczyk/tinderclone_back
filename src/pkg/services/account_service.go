package services

import "tinderclone_back/src/pkg/stores"

type IAccounter interface {
	GetAccountInformations(username string) Result
}

type accounter struct {
}

func NewAccounter() *accounter {
	return &accounter{}
}

func (s *accounter) GetAccountInformations(username string) Result {
	user, err := stores.SelectUserByUsername(username)
	if err != nil {
		return *CreateServiceResult("Couldn't find an user", 404, []interface{}{})
	}

	return *CreateServiceResult("User found", 200, []interface{}{user})
}
