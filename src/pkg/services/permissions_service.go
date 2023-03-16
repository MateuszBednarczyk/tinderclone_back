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
	err := s.userStore.UpdateUserRole(username, domain.Role(2))
	if err != nil {
		return CreateServiceResult(err.Error(), 500, []interface{}{})
	}

	return CreateServiceResult("Permission updated", 200, []interface{}{})
}
