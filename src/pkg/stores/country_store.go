package stores

import (
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type ICountryStore interface {
	SelectCountryByName(countryName string) *domain.Country
	IsCountryAlreadyAvailable(countryName string) bool
	SaveCountry(entity *domain.Country) error
}

type countryStore struct {
	db *gorm.DB
}

func NewCountryStore(db *gorm.DB) *countryStore {
	return &countryStore{
		db: db,
	}
}

func (s *countryStore) SelectCountryByName(countryName string) *domain.Country {
	var country domain.Country
	err := s.db.Where("country_name = ?", countryName).Find(&country)
	if err.Error != nil {
		return nil
	}

	if country.CountryName == "" {
		return nil
	}

	return &country
}

func (s *countryStore) IsCountryAlreadyAvailable(countryName string) bool {
	var country domain.Country
	_ = s.db.Where("country_name = ?", countryName).Find(&country)

	return country.CountryName != ""
}

func (s *countryStore) SaveCountry(entity *domain.Country) error {
	result := s.db.Create(&entity)
	return result.Error
}
