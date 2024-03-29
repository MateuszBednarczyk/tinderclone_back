package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/services"
	"tinderclone_back/src/pkg/utils"
	"tinderclone_back/src/test/mocks"
)

func TestGetAccountInformations(t *testing.T) {
	mockUser := domain.User{
		UserID:    uuid.New(),
		Username:  "admin",
		Password:  "pass",
		Name:      "Admin",
		Surname:   "Adminsky",
		Countries: []domain.Country{},
		Cities:    []domain.City{},
		Role:      domain.Role(2),
	}

	userStoreMock := mocks.IUserStore{}
	userUtil := utils.NewUserUtil()
	userStoreMock.On("SelectUserByUsername", mock.Anything).Return(&mockUser, nil)

	serviceInstance := services.NewAccounter(&userStoreMock, userUtil)
	result := serviceInstance.GetAccountInformations("admin")

	require.Equal(t, mockUser.Username, result.Content[0].(*dto.User).Username)
}
