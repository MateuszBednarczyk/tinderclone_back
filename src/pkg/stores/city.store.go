package stores

import (
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type ICityStore interface {
	SelectCityByName(cityName string) *domain.City
	IsCityAlreadyAvailable(cityName string) bool
	SaveNewCity(entity *domain.City) error
}

type cityStore struct {
	db *gorm.DB
}

func NewCityStore(db *gorm.DB) *cityStore {
	return &cityStore{
		db: db,
	}
}

func (s *cityStore) SelectCityByName(cityName string) *domain.City {
	var city domain.City
	err := s.db.Where("city_name = ?", cityName).Find(&city)
	if err.Error != nil {
		return nil
	}

	if city.CityName == "" {
		return nil
	}

	return &city
}

func (s *cityStore) SaveNewCity(entity *domain.City) error {
	result := s.db.Create(&entity)
	return result.Error
}

func (s *cityStore) IsCityAlreadyAvailable(cityName string) bool {
	var city domain.City
	_ = s.db.Where("city_name = ?", cityName).Find(&city)

	return city.CityName != ""
}
