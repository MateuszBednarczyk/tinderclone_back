package test

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/services"
	"tinderclone_back/src/test/mocks"
)

func TestSaveCity(t *testing.T) {
	countryEntityMock := &domain.Country{
		CountryID:   uuid.New(),
		CountryName: "POLAND",
	}

	countryStoreMock := mocks.ICountryStore{}
	countryStoreMock.On("IsCountryAlreadyAvailable", mock.Anything).Return(true)
	countryStoreMock.On("SelectCountryByName", mock.Anything).Return(countryEntityMock, nil)

	cityStoreMock := mocks.ICityStore{}
	cityStoreMock.On("SaveNewCity", mock.Anything).Return(nil)
	cityStoreMock.On("IsCityInCountryAlreadyAvailable", mock.Anything, mock.Anything).Return(false)

	serviceInstance := services.NewCitier(&cityStoreMock, &countryStoreMock)
	result := serviceInstance.SaveNewCity("warsaw", "POLAND")

	log.Print(result.Message)
	require.Equal(t, 200, result.Code)
}
