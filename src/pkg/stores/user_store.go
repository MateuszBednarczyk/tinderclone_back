package stores

import (
	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/domain"
)

type IUserStore interface {
	SaveUser(entity *domain.User) error
	IsUsernameAlreadyTaken(username string) bool
	SelectUserByUsername(username string) (*domain.User, error)
}

type userStore struct {
}

func NewUserStore() *userStore {
	return &userStore{}
}

func (s *userStore) SaveUser(entity *domain.User) error {
	result := database.GetDb().Create(&entity)
	return result.Error
}

func (s *userStore) IsUsernameAlreadyTaken(username string) bool {
	var user domain.User
	_ = database.GetDb().Where("username = ?", username).Find(&user)

	return user.Username != ""
}

func (s *userStore) SelectUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := database.GetDb().Model(domain.User{Username: username}).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}
