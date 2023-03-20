package services

import (
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

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

	countries := []string{}
	cities := []string{}

	for _, country := range user.Countries {
		countries = append(countries, country.CountryName)
	}

	for _, city := range user.Cities {
		cities = append(cities, city.CityName)
	}

	response := dto.User{
		Id:        user.UserID,
		Username:  user.Username,
		Name:      user.Name,
		Surname:   user.Surname,
		Countries: countries,
		Cities:    cities,
	}

	return CreateServiceResult("User found", 200, []interface{}{response})
}
