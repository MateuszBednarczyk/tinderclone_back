package stores

import (
	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/domain"
)

func SelectCityByName(cityName string) *domain.City {
	var city domain.City
	err := database.GetDb().Where("city_name = ?", cityName).Find(&city)
	if err.Error != nil {
		return nil
	}

	if city.CityName == "" {
		return nil
	}

	return &city
}

func IsCityAlreadyAvailable(cityName string) bool {
	var city domain.City
	_ = database.GetDb().Where("city_name = ?", cityName).Find(&city)

	return city.CityName != ""
}
