package utils

import (
	"tinderclone_back/src/pkg/domain"
	"tinderclone_back/src/pkg/dto"
)

type IUserUtil interface {
	ProcessUserEntityToDTO(user *domain.User) *dto.User
	processCitiesToStringSlice(user *domain.User) []string
	processCountriesToStringSlice(user *domain.User) []string
}

type userUtil struct {
}

func NewUserUtil() *userUtil {
	return &userUtil{}
}

func (u *userUtil) ProcessUserEntityToDTO(user *domain.User) *dto.User {
	return &dto.User{
		Id:        user.UserID,
		Username:  user.Username,
		Name:      user.Name,
		Countries: u.processCountriesToStringSlice(user),
		Cities:    u.processCitiesToStringSlice(user),
		Role:      user.Role,
	}
}

func (u *userUtil) processCitiesToStringSlice(user *domain.User) []string {
	result := []string{}
	for _, city := range user.Cities {
		result = append(result, city.CityName)
	}

	return result
}

func (u *userUtil) processCountriesToStringSlice(user *domain.User) []string {
	result := []string{}
	for _, country := range user.Countries {
		result = append(result, country.CountryName)
	}

	return result
}
