package services

import (
	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/stores"
)

type ICountrier interface {
	SaveNewCountry(countryName string) *Result
}

type countrier struct {
}

func NewCountrier() *countrier {
	return &countrier{}
}

func (s *countrier) SaveNewCountry(countryName string) *Result {
	if stores.IsCountryAlreadyAvailable(countryName) {
		return CreateServiceResult("Country is already available", 409, []interface{}{})
	}

	err := stores.SaveCountry(&domain.Country{
		CountryName: countryName,
		Cities:      []domain.City{},
	})
	if err != nil {
		return CreateServiceResult("Couldn't save a country", 500, []interface{}{})
	}

	return CreateServiceResult("Country saved", 200, []interface{}{})
}
