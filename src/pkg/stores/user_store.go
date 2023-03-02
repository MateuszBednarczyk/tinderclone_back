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
