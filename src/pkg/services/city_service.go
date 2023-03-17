package services

import (
	"strings"

	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/stores"
)

type ICitier interface {
	SaveNewCity(cityName string, countryName string) *Result
	GetAllAvailableCitiesInGivenCountry(countryName string, cityName string) *Result
}

type citier struct {
	cityStore    stores.ICityStore
	countryStore stores.ICountryStore
}

func NewCitier(cityStore stores.ICityStore, countryStore stores.ICountryStore) *citier {
	return &citier{cityStore: cityStore, countryStore: countryStore}
}

func (s *citier) SaveNewCity(cityName string, countryName string) *Result {
	var err error
	cityName = strings.ToUpper(cityName)
	countryName = strings.ToUpper(countryName)

	if !s.countryStore.IsCountryAlreadyAvailable(countryName) {
		return CreateServiceResult("Can't save city, because country doesn't persist", 404, []interface{}{})
	}

	country, err := s.countryStore.SelectCountryByName(countryName)
	if s.cityStore.IsCityInCountryAlreadyAvailable(cityName, country.CountryID) {
		return CreateServiceResult("City is already available", 409, []interface{}{})
	}
	if err != nil {
		return CreateServiceResult("Error while trying to get the city from db", 500, []interface{}{})
	}

	city := &domain.City{
		CityName: strings.ToUpper(cityName),
		Country:  *country,
	}
	err = s.cityStore.SaveNewCity(city)

	if err != nil {
		return CreateServiceResult("Couldn't save city", 500, []interface{}{})
	}
	country.Cities = append(country.Cities, *city)

	return CreateServiceResult("City has been saved", 200, []interface{}{})
}

func (s *citier) GetAllAvailableCitiesInGivenCountry(countryName string, cityName string) *Result {
	availableCities := []string{}
	country, err := s.countryStore.SelectCountryByName(countryName)
	if err != nil {
		return CreateServiceResult("Country does not exist", 404, []interface{}{})
	}
	for _, city := range country.Cities {
		availableCities = append(availableCities, city.CityName)
	}
	return CreateServiceResult("List of available cities in"+countryName, 200, []interface{}{availableCities})
}
