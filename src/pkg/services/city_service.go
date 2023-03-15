package services

import (
	"strings"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/stores"
)

type ICitier interface {
	SaveNewCity(cityName string, countryName string) *Result
}

type citier struct {
	cityStore    stores.ICityStore
	countryStore stores.ICountryStore
}

func NewCitier(cityStore stores.ICityStore, countryStore stores.ICountryStore) *citier {
	return &citier{}
}

func (s *citier) SaveNewCity(cityName string, countryName string) *Result {
	if s.countryStore.IsCityAlreadyAvailableInCountry(countryName, cityName) {
		return CreateServiceResult("City is already available", 401, []interface{}{})
	}
	country := s.countryStore.SelectCountryByName(countryName)
	city := &domain.City{
		CityName: strings.ToUpper(cityName),
		Country:  *country,
	}
	err := s.cityStore.SaveNewCity(city)

	if err != nil {
		return CreateServiceResult("Couldn't save city", 500, []interface{}{})
	}
	country.Cities = append(country.Cities, *city)

	return CreateServiceResult("City has been saved", 200, []interface{}{})
}
