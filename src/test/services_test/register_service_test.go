package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/services"
	"tinderclone_back/src/test/mocks"
)

func TestHashPassword(t *testing.T) {
	plainPass := "ADMIN"
	result, err := services.HashPassword(plainPass)

	require.NotEqual(t, result, nil)
	require.Equal(t, err, nil)
}

func TestIsPasswordValid(t *testing.T) {
	pass := ""
	result := services.IsPasswordValid(pass)

	require.Equal(t, result, false)
}

func TestIsUsernameValid(t *testing.T) {
	u1 := ""
	u2 := "AD"

	r1 := services.IsUsernameValid(u1)
	r2 := services.IsUsernameValid(u2)

	require.Equal(t, r1, false)
	require.Equal(t, r2, true)
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

	city := domain.City{
		CityID:    uuid.New(),
		CityName:  "WWA",
		CountryID: country.CountryID,
		Country:   country,
	}

	userStoreMock := mocks.IUserStore{}
	userStoreMock.On("SaveUser", mock.Anything).Return(nil)
	userStoreMock.On("IsUsernameAlreadyTaken", mock.Anything).Return(false)

	countryStoreMock := mocks.ICountryStore{}
	countryStoreMock.On("IsCountryAlreadyAvailable").Return(true)
	countryStoreMock.On("SelectCountryByName", mock.Anything).Return(&country, nil)

	cityStoreMock := mocks.ICityStore{}
	cityStoreMock.On("IsCityAlreadyAvailable", mock.Anything).Return(true)
	cityStoreMock.On("SelectCityByName", mock.Anything).Return(&city)

	serviceInstance := services.NewAccountMaker(&userStoreMock, &countryStoreMock, &cityStoreMock)
	result := serviceInstance.RegisterUser(requestBody)

	require.Equal(t, result.Message, "Account created")
}
