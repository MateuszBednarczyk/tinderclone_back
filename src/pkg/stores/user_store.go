package stores

import (
	"gorm.io/gorm"

	"tinderclone_back/src/pkg/domain"
)

type IUserStore interface {
	SaveUser(entity *domain.User) error
	IsUsernameAlreadyTaken(username string) bool
	SelectUserByUsername(username string) (*domain.User, error)
	UpdateUserRole(username string, role domain.Role) error
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
	err := s.db.First(&user, "username = ?", username)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (s *userStore) UpdateUserRole(username string, role domain.Role) error {
	result := s.db.Model(&domain.User{}).Where("username = ?", username).Update("role", role)
	return result.Error
}
