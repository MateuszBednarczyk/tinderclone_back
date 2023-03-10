package stores

import (
	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/domain"
)

func SelectCountryByName(countryName string) *domain.Country {
	var country domain.Country
	err := database.GetDb().Where("country_name = ?", countryName).Find(&country)
	if err.Error != nil {
		return nil
	}

	if country.CountryName == "" {
		return nil
	}

	return &country
}

func IsCountryAlreadyAvailable(countryName string) bool {
	var country domain.Country
	_ = database.GetDb().Where("country_name = ?", countryName).Find(&country)

	return country.CountryName != ""
}

func SaveCountry(entity *domain.Country) error {
	result := database.GetDb().Create(&entity)
	return result.Error
}
