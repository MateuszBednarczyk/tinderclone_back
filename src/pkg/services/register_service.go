package services

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/stores"
)

type IAccountMaker interface {
	RegisterUser(dto dto.RegisterUser) *Result
}

type accountMaker struct {
	userStore    stores.IUserStore
	countryStore stores.ICountryStore
	cityStore    stores.ICityStore
}

func NewAccountMaker(userStore stores.IUserStore, countryStore stores.ICountryStore, cityStore stores.ICityStore) *accountMaker {
	return &accountMaker{
		userStore:    userStore,
		countryStore: countryStore,
		cityStore:    cityStore,
	}
}

func (s *accountMaker) RegisterUser(requestBody dto.RegisterUser) *Result {
	if !IsUsernameValid(requestBody.Username) {
		return CreateServiceResult("Username cannot be blank", 400, []interface{}{})
	}

	if !IsPasswordValid(requestBody.Password) {
		return CreateServiceResult("Password cannot be blank", 400, []interface{}{})
	}

	if s.userStore.IsUsernameAlreadyTaken(requestBody.Username) {
		return CreateServiceResult("Username is already taken", 409, []interface{}{})
	}

	hash, err := HashPassword(requestBody.Password)
	if err != nil {
		return CreateServiceResult("Couldn't hash password", 500, []interface{}{})
	}

	cities := []domain.City{}
	for _, city := range requestBody.Cities {
		cityEntity := s.cityStore.SelectCityByName(city)
		if cityEntity == nil {
			return CreateServiceResult("Service is not available in: "+city, 404, []interface{}{})
		}
		cities = append(cities, *cityEntity)
	}

	countries := []domain.Country{}
	for _, country := range requestBody.Countries {
		countryEntity, err := s.countryStore.SelectCountryByName(country)
		if countryEntity == nil {
			return CreateServiceResult("Service is not available in: "+country, 404, []interface{}{})
		}
		if err != nil {
			return CreateServiceResult("Error while trying to get country from db", 500, []interface{}{})
		}
		countries = append(countries, *countryEntity)
	}

	user := &domain.User{
		Username:  requestBody.Username,
		Password:  string(hash),
		Name:      requestBody.Name,
		Surname:   requestBody.Surname,
		Countries: countries,
		Cities:    cities,
		Role:      domain.Role(2),
	}

	result := s.userStore.SaveUser(user)
	if result != nil {
		return CreateServiceResult("There was an error, while attempt of saving user", 500, []interface{}{})
	}

	return CreateServiceResult("Account created", 200, []interface{}{})
}

func IsUsernameValid(username string) bool {
	return len(strings.TrimSpace(username)) > 0
}

func IsPasswordValid(password string) bool {
	return len(strings.TrimSpace(password)) > 0
}

func HashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
