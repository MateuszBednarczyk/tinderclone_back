package services

import (
	"strings"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/stores"
)

type ICountrier interface {
	SaveNewCountry(countryName string) *Result
	GetAllCountriesNames() *Result
}

type countrier struct {
	countryStore stores.ICountryStore
}

func NewCountrier(store stores.ICountryStore) *countrier {
	return &countrier{
		countryStore: store,
	}
}

func (s *countrier) SaveNewCountry(countryName string) *Result {
	if s.countryStore.IsCountryAlreadyAvailable(countryName) {
		return CreateServiceResult("Country is already available", 409, []interface{}{})
	}

	err := s.countryStore.SaveCountry(&domain.Country{
		CountryName: strings.ToUpper(countryName),
	})
	if err != nil {
		return CreateServiceResult("Couldn't save a country", 500, []interface{}{})
	}

	return CreateServiceResult("Country saved", 200, []interface{}{})
}

func (s *countrier) GetAllCountriesNames() *Result {
	countries, err := s.countryStore.GetAllCountriesNames()
	if err != nil {
		return CreateServiceResult("Couldn't find countries", 404, []interface{}{})
	}

	return CreateServiceResult("List of available countries", 200, []interface{}{countries})
}
