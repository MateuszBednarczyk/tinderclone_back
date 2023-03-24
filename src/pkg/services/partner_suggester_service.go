package services

import (
	"github.com/google/uuid"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
	"tinderclone_back/src/pkg/utils"
)

type IPartnerSuggester interface {
	SuggestPartners(userDTO dto.User) *Result
}

type partnerSuggester struct {
	userStore    stores.IUserStore
	cityStore    stores.ICityStore
	countryStore stores.ICountryStore
	userUtil     utils.IUserUtil
}

func NewPartnerSuggester(userStore stores.IUserStore, cityStore stores.ICityStore, countryStore stores.ICountryStore, userUtil utils.IUserUtil) *partnerSuggester {
	return &partnerSuggester{
		userStore:    userStore,
		cityStore:    cityStore,
		countryStore: countryStore,
		userUtil:     userUtil,
	}
}

func (s *partnerSuggester) SuggestPartners(userDTO dto.User) *Result {
	userEntity, err := s.userStore.SelectUserByUsername(userDTO.Username)
	if err != nil {
		return CreateServiceResult("Cannot find user", 404, []interface{}{})
	}

	usersSuggestions := []domain.User{}
	for _, city := range userEntity.Cities {
		users, err := s.userStore.GetAllUsersFromGivenCity(city.CityID)
		if err != nil {
			return CreateServiceResult("Cannot process users", 500, []interface{}{})
		}
		for _, user := range users {
			if !containsUser(usersSuggestions, user.UserID) && !containsUser(userEntity.DislikedUsers, user.UserID) {
				usersSuggestions = append(usersSuggestions, user)
			}
		}
	}

	for _, country := range userEntity.Countries {
		users, err := s.userStore.GetAllUsersFromGivenCountry(country.CountryID)
		if err != nil {
			return CreateServiceResult("Cannot process users", 500, []interface{}{})
		}
		for _, user := range users {
			if !containsUser(usersSuggestions, user.UserID) && !containsUser(userEntity.DislikedUsers, user.UserID) {
				usersSuggestions = append(usersSuggestions, user)
			}
		}
	}

	result := []dto.User{}
	for _, entity := range usersSuggestions {
		processedDTO := s.userUtil.ProcessUserEntityToDTO(&entity)
		result = append(result, *processedDTO)
	}

	return CreateServiceResult("ok", 200, []interface{}{result})
}

func containsUser(slice []domain.User, id uuid.UUID) bool {
	for _, user := range slice {
		if user.UserID == id {
			return true
		}
	}
	return false
}
