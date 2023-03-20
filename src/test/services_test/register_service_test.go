package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/services"
	"tinderclone_back/src/test/mocks"
)

func TestHashPassword(t *testing.T) {
	//given
	plainPass := "ADMIN"

	//when
	result, err := services.HashPassword(plainPass)

	//then
	assert.NotEqual(t, result, nil)
	assert.Equal(t, err, nil)
}

func TestIsPasswordValid(t *testing.T) {
	//given
	pass := ""

	//when
	result := services.IsPasswordValid(pass)

	//then
	assert.Equal(t, result, false)
}

func TestIsUsernameValid(t *testing.T) {
	//given
	u1 := ""
	u2 := "AD"

	//when
	r1 := services.IsUsernameValid(u1)
	r2 := services.IsUsernameValid(u2)

	//then
	assert.Equal(t, r1, false)
	assert.Equal(t, r2, true)
}

func TestRegisterUser(t *testing.T) {
	requestBody := dto.RegisterUser{
		Username:  "a",
		Password:  "b",
		Name:      "c",
		Surname:   "d",
		Countries: []string{"PL"},
		Cities:    []string{"WWA"},
	}

	country := domain.Country{
		CountryID:   uuid.New(),
		CountryName: "PL",
	}

	userStoreMock := mocks.IUserStore{}
	userStoreMock.On("SaveUser", mock.Anything).Return(nil)
	userStoreMock.On("IsUsernameAlreadyTaken", mock.Anything).Return(false)

	cityStoreMock := mocks.ICityStore{}
	cityStoreMock.On("IsCityAlreadyAvailable", mock.Anything).Return(true)
	cityStoreMock.On("SelectCityByName", mock.Anything).Return(&domain.City{
		CityID:   uuid.New(),
		CityName: "WWA",
		Country:  country,
	})

	countryStoreMock := mocks.ICountryStore{}
	countryStoreMock.On("IsCountryAlreadyAvailable").Return(true)
	countryStoreMock.On("SelectCountryByName", mock.Anything).Return(&country)

	serviceInstance := services.NewAccountMaker(&userStoreMock, &countryStoreMock, &cityStoreMock)
	result := serviceInstance.RegisterUser(requestBody)

	assert.Equal(t, 200, result.Code)
}
