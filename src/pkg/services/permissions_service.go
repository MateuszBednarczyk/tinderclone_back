package services

import (
	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/stores"
)

type IPermitter interface {
	GiveUserAdminPermission(username string) *Result
}

type permitter struct {
	userStore stores.IUserStore
}

func NewPermitter(userStore stores.IUserStore) *permitter {
	return &permitter{
		userStore: userStore,
	}
}

func (s *permitter) GiveUserAdminPermission(username string) *Result {
	var err error
	user, err := s.userStore.SelectUserByUsername(username)
	if err != nil {
		return CreateServiceResult("User with given username does not exist", 404, []interface{}{})
	}
	if user.Role == domain.Role(2) {
		return CreateServiceResult("User is already admin", 409, []interface{}{})
	}

	err = s.userStore.UpdateUserRole(username, domain.Role(2))
	if err != nil {
		return CreateServiceResult(err.Error(), 500, []interface{}{})
	}

	return CreateServiceResult("Permission updated", 200, []interface{}{})
}
