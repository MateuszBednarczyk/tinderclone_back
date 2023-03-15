package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/services"
	"tinderclone_back/src/test/mocks"
)

func TestGetAccountInformations(t *testing.T) {
	mockUser := domain.User{
		UserID:   uuid.New(),
		Username: "admin",
		Password: "pass",
		Name:     "Admin",
		Surname:  "Adminsky",
		Country:  domain.Country{},
		City:     domain.City{},
		Role:     domain.Role(2),
	}

	userStoreMock := mocks.IUserStore{}
	userStoreMock.On("SelectUserByUsername", mock.Anything).Return(&mockUser, nil)

	serviceInstance := services.NewAccounter(&userStoreMock)
	result := serviceInstance.GetAccountInformations("admin")

	assert.Equal(t, mockUser.Username, result.Content[0].(*domain.User).Username)
}
