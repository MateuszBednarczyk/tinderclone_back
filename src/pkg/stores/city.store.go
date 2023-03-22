package stores

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type ICityStore interface {
	SelectCityByName(cityName string) *domain.City
	IsCityAlreadyAvailable(cityName string) bool
	SaveNewCity(entity *domain.City) error
	IsCityInCountryAlreadyAvailable(cityName string, countryID uuid.UUID) bool
	SelectAllCitiesWhereCountryIdEqual(countryID uuid.UUID) ([]domain.City, error)
}

type cityStore struct {
	db *gorm.DB
}

func NewCityStore(db *gorm.DB) *cityStore {
	return &cityStore{
		db: db,
	}
}

func (s *cityStore) SelectAllCitiesWhereCountryIdEqual(countryID uuid.UUID) ([]domain.City, error) {
	cities := []domain.City{}
	err := s.db.Find(&cities).Where("country_id = ?", countryID)
	if err != nil {
		return nil, err.Error
	}
	return cities, nil
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

func (s *cityStore) IsCityInCountryAlreadyAvailable(cityName string, countryID uuid.UUID) bool {
	var city domain.City
	_ = s.db.Table("cities").Where("country_id = ? AND city_name = ?", countryID, cityName).Find(&city)
	log.Print(city.CityName, city.Country)

	return city.CityName != ""
}
