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
}

func NewAccountMaker() *accountMaker {
	return &accountMaker{}
}

func (s *accountMaker) RegisterUser(requestBody dto.RegisterUser) *Result {
	if !isUsernameValid(requestBody.Username) {
		return CreateServiceResult("Username cannot be blank", 400, []interface{}{})
	}

	if !isPasswordValid(requestBody.Password) {
		return CreateServiceResult("Password cannot be blank", 400, []interface{}{})
	}

	if stores.IsUsernameAlreadyTaken(requestBody.Username) {
		return CreateServiceResult("Username is already taken", 409, []interface{}{})
	}

	hash, err := hashPassword(requestBody.Password)
	if err != nil {
		return CreateServiceResult("Couldn't hash password", 500, []interface{}{})
	}

	city := stores.SelectCityByName(requestBody.CityName)
	if city == nil {
		return CreateServiceResult("Service is not available in your city", 403, []interface{}{})
	}

	country := stores.SelectCountryByName(requestBody.Country)
	if country == nil {
		return CreateServiceResult("Service is not available in your country", 403, []interface{}{})
	}

	user := &domain.User{
		Username: requestBody.Username,
		Password: string(hash),
		Name:     requestBody.Name,
		Surname:  requestBody.Surname,
		Country:  country.CountryID,
		City:     city.CityID,
		Role:     2,
	}

	result := stores.SaveUser(user)
	if result != nil {
		return CreateServiceResult("There was an error, while attempt of saving user", 500, []interface{}{})
	}

	return CreateServiceResult("Account created", 200, []interface{}{})
}

func isUsernameValid(username string) bool {
	return len(strings.TrimSpace(username)) > 0
}

func isPasswordValid(password string) bool {
	return len(strings.TrimSpace(password)) > 0
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
