package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type IUserStore interface {
	SaveUser(entity *domain.User) error
	IsUsernameAlreadyTaken(username string) bool
	SelectUserByUsername(username string) (*domain.User, error)
	UpdateUserRole(username string, role domain.Role) error
	GetAllUsersFromGivenCountry(countryID uuid.UUID) ([]domain.User, error)
	GetAllUsersFromGivenCity(cityID uuid.UUID) ([]domain.User, error)
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		db: db,
	}
}

func (s *userStore) SaveUser(entity *domain.User) error {
	result := s.db.Create(&entity)
	return result.Error
}

func (s *userStore) IsUsernameAlreadyTaken(username string) bool {
	var user domain.User
	_ = s.db.Where("username = ?", username).Find(&user)

	return user.Username != ""
}

func (s *userStore) SelectUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := s.db.First(&user, "username = ?", username).Preload("Cities").Preload("Countries").First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (s *userStore) UpdateUserRole(username string, role domain.Role) error {
	result := s.db.Model(&domain.User{}).Where("username = ?", username).Update("role", role)
	return result.Error
}

func (s *userStore) GetAllUsersFromGivenCountry(countryID uuid.UUID) ([]domain.User, error) {
	var userIDs []uuid.UUID
	var entities []domain.User
	err := s.db.Select("user_user_id").Table("users_countries").Where("country_country_id = ?", countryID).Find(&userIDs).Error
	if err != nil {
		return nil, err
	}

	for _, userID := range userIDs {
		var entity domain.User
		err := s.db.Where("user_id = ?", userID).Find(&entity).Error
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (s *userStore) GetAllUsersFromGivenCity(cityID uuid.UUID) ([]domain.User, error) {
	var userIDs []uuid.UUID
	var entities []domain.User
	err := s.db.Select("user_user_id").Table("users_cities").Where("city_city_id = ?", cityID).Find(&userIDs).Error
	if err != nil {
		return nil, err
	}

	for _, userID := range userIDs {
		var entity domain.User
		err := s.db.Where("user_id = ?", userID).Preload("Cities").Preload("Countries").First(&entity).Error
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}
