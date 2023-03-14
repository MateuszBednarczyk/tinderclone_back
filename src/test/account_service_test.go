package test

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"tinderclone_back/src/pkg/services"
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

// func TestRegisterUser(t *testing.T) {
// 	requestBody := dto.RegisterUser{
// 		Username: "a",
// 		Password: "b",
// 		Name:     "c",
// 		Surname:  "d",
// 		Country:  "PL",
// 		CityName: "WWA",
// 	}

// 	userStoreMock := mocks.IUserStore{}
// 	userStoreMock.On("SaveUser", mock.Anything).Return(nil)
// 	userStoreMock.On("IsUsernameAlreadyTaken", mock.Anything).Return(false)

// 	serviceInstance := services.NewAccountMaker(&userStoreMock, stores.NewCountryStore(), stores.NewCityStore())
// 	result := serviceInstance.RegisterUser(requestBody)

// 	assert.Equal(t, 200, result.Code)
// }
