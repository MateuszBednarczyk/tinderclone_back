package stores

import (
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type ICountryStore interface {
	SelectCountryByName(countryName string) (*domain.Country, error)
	IsCountryAlreadyAvailable(countryName string) bool
	SaveCountry(entity *domain.Country) error
	GetAllCountriesNames() ([]string, error)
}

type countryStore struct {
	db *gorm.DB
}

func NewCountryStore(db *gorm.DB) *countryStore {
	return &countryStore{
		db: db,
	}
}

func (s *countryStore) SelectCountryByName(countryName string) (*domain.Country, error) {
	var country domain.Country
	err := s.db.Where("country_name = ?", countryName).Find(&country)
	if err.Error != nil {
		return nil, err.Error
	}

	if country.CountryName == "" {
		return nil, err.Error
	}

	return &country, nil
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

func (s *countryStore) GetAllCountriesNames() ([]string, error) {
	var countries []string
	err := s.db.Select("country_name").Find(&countries)

	if err != nil {
		return nil, err.Error
	}

	return countries, nil
}
