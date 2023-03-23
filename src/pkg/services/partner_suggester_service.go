package services

import (
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

type IPartnerSuggester interface {
	SuggestPartners(userDTO dto.User) *Result
}

type partnerSuggester struct {
	userStore    stores.IUserStore
	cityStore    stores.ICityStore
	countryStore stores.ICountryStore
}

func NewPartnerSuggester(userStore stores.IUserStore, cityStore stores.ICityStore, countryStore stores.ICountryStore) *partnerSuggester {
	return &partnerSuggester{
		userStore:    userStore,
		cityStore:    cityStore,
		countryStore: countryStore,
	}
}

func (s *partnerSuggester) SuggestPartners(userDTO dto.User) *Result {
	_, err := s.userStore.SelectUserByUsername(userDTO.Username)
	if err != nil {
		return CreateServiceResult("Cannot find user", 404, []interface{}{})
	}
	return CreateServiceResult("ok", 200, []interface{}{})
}
