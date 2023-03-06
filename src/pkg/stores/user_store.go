package stores

import (
	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/domain"
)

func SaveUser(entity *domain.User) error {
	result := database.GetDb().Create(&entity)
	return result.Error
}

func IsUsernameAlreadyTaken(username string) bool {
	var user domain.User
	_ = database.GetDb().Where("username = ?", username).Find(&user)

	return user.Username != ""
}

func SelectUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := database.GetDb().Where("username = ?", username).Find(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}
